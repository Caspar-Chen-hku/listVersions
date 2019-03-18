package main

import (
	"context"
	"fmt"
	"sort"

	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"
)

// same determines whether the two versions are in the same release, that is, they
// may only deffer in batch
func same(v1 *semver.Version,v2 *semver.Version) bool {
	return v1.Major==v2.Major && v1.Minor==v2.Minor
}

// LatestVersions returns a sorted slice with the highest version as its first element and the highest version of the smaller minor versions in a descending order
func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	sort.Slice(releases, func(i, j int) bool {
		return !(releases[i].LessThan(*releases[j]))
	})

	var versionSlice []*semver.Version
	var num_releases = len(releases)

	for i := 0; i < num_releases; i++ {
		if releases[i].LessThan(*minVersion)||(i!=0&&same(releases[i],releases[i-1])){
			continue
		}
     	versionSlice = append(versionSlice, releases[i])     	
   	}
	
	// This is just an example structure of the code, if you implement this interface, the test cases in main_test.go are very easy to run
	return versionSlice
}

// Here we implement the basics of communicating with github through the library as well as printing the version
// You will need to implement LatestVersions function as well as make this application support the file format outlined in the README
// Please use the format defined by the fmt.Printf line at the bottom, as we will define a passing coding challenge as one that outputs
// the correct information, including this line
func main() {
	// Github
	client := github.NewClient(nil)
	ctx := context.Background()
	opt := &github.ListOptions{PerPage: 10}
	releases, _, err := client.Repositories.ListReleases(ctx, "kubernetes", "kubernetes", opt)
	if err != nil {
		//panic should only be used when some fatal error happens, and everything
		//should crash immediately. It's better to avoid using panic
		fmt.Print(err)
		//panic(err) // is this really a good way?
	}
	minVersion := semver.New("1.8.0")
	allReleases := make([]*semver.Version, len(releases))
	for i, release := range releases {
		versionString := *release.TagName
		if versionString[0] == 'v' {
			versionString = versionString[1:]
		}
		allReleases[i] = semver.New(versionString)
	}
	versionSlice := LatestVersions(allReleases, minVersion)

	fmt.Printf("latest versions of kubernetes/kubernetes: %s", versionSlice)
}
