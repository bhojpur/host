package amazonec2

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"errors"
)

type region struct {
	AmiId string
}

// Ubuntu 20.04 LTS hvm:ebs-ssd (amd64)
// See https://cloud-images.ubuntu.com/locator/ec2/
var regionDetails map[string]*region = map[string]*region{
	"af-south-1":      {"ami-0f974a103be1a32cc"},
	"ap-northeast-1":  {"ami-09ff2b6ef00accc2e"},
	"ap-northeast-2":  {"ami-0b329fb1f17558744"},
	"ap-northeast-3":  {"ami-0fa06c85437a220f5"},
	"ap-southeast-1":  {"ami-048b4b1ddefe6759f"},
	"ap-southeast-2":  {"ami-052a251c7ca533c26"},
	"ap-south-1":      {"ami-01957c76cce45de38"},
	"ap-east-1":       {"ami-0a7893426ebe74a8c"},
	"ca-central-1":    {"ami-095509bf36d02a8e0"},
	"cn-north-1":      {"ami-0592ccadb56e65f8d"},
	"cn-northwest-1":  {"ami-007d0f254ea0f8588"},
	"eu-central-1":    {"ami-0d3905203a039e3b0"},
	"eu-north-1":      {"ami-0803bdfa576dd60c2"},
	"eu-south-1":      {"ami-08a4feb1e48d2f9d8"},
	"eu-west-1":       {"ami-0b7fd7bc9c6fb1c78"},
	"eu-west-2":       {"ami-02ead6ecbd926d792"},
	"eu-west-3":       {"ami-0d7b738ade930e24a"},
	"me-south-1":      {"ami-0e53ffff39fa7b5c1"},
	"sa-east-1":       {"ami-03f2389c2526e67bd"},
	"us-east-1":       {"ami-04cc2b0ad9e30a9c8"},
	"us-east-2":       {"ami-02fc6052104add5ae"},
	"us-west-1":       {"ami-07be40433001d2433"},
	"us-west-2":       {"ami-0a62a78cfedc09d76"},
	"us-gov-west-1":   {"ami-0c39aacd1cc8a1ccf"},
	"us-gov-east-1":   {"ami-0dec4096f1af85e9b"},
	"custom-endpoint": {""},
}

func awsRegionsList() []string {
	var list []string

	for k := range regionDetails {
		list = append(list, k)
	}

	return list
}

func validateAwsRegion(region string) (string, error) {
	for _, v := range awsRegionsList() {
		if v == region {
			return region, nil
		}
	}

	return "", errors.New("Invalid region specified")
}
