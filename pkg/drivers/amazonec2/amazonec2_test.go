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
	"testing"

	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	ctest "github.com/bhojpur/host/cmd/machine/commands/test"
	"github.com/bhojpur/host/pkg/version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	testSSHPort     = int64(22)
	testBhojpurPort = int64(2376)
	testSwarmPort   = int64(3376)
)

var (
	/* This test resource should be used in tests that set their own IpPermissions */
	securityGroup = &ec2.SecurityGroup{
		GroupName: aws.String("test-group"),
		GroupId:   aws.String("12345"),
		VpcId:     aws.String("12345"),
	}

	/* This test resource should only be used in tests that do not update IpPermissions */
	securityGroupNoIpPermissions = &ec2.SecurityGroup{
		GroupName:     aws.String("test-group"),
		GroupId:       aws.String("12345"),
		VpcId:         aws.String("12345"),
		IpPermissions: []*ec2.IpPermission{},
	}
)

func TestConfigureSecurityGroupPermissionsEmpty(t *testing.T) {
	driver := NewTestDriver()

	perms, err := driver.configureSecurityGroupPermissions(securityGroup)

	assert.Nil(t, err)
	assert.Len(t, perms, 2)
}

func TestConfigureSecurityGroupPermissionsSshOnly(t *testing.T) {
	driver := NewTestDriver()
	group := securityGroup
	group.IpPermissions = []*ec2.IpPermission{
		{
			IpProtocol: aws.String("tcp"),
			FromPort:   aws.Int64(int64(testSSHPort)),
			ToPort:     aws.Int64(int64(testSSHPort)),
		},
	}

	perms, err := driver.configureSecurityGroupPermissions(group)

	assert.Nil(t, err)
	assert.Len(t, perms, 1)
	assert.Equal(t, testBhojpurPort, *perms[0].FromPort)
}

func TestConfigureSecurityGroupPermissionsBhojpurOnly(t *testing.T) {
	driver := NewTestDriver()
	group := securityGroup
	group.IpPermissions = []*ec2.IpPermission{
		{
			IpProtocol: aws.String("tcp"),
			FromPort:   aws.Int64((testBhojpurPort)),
			ToPort:     aws.Int64((testBhojpurPort)),
		},
	}

	perms, err := driver.configureSecurityGroupPermissions(group)

	assert.Nil(t, err)
	assert.Len(t, perms, 1)
	assert.Equal(t, testSSHPort, *perms[0].FromPort)
}

func TestConfigureSecurityGroupPermissionsBhojpurAndSsh(t *testing.T) {
	driver := NewTestDriver()
	group := securityGroup
	group.IpPermissions = []*ec2.IpPermission{
		{
			IpProtocol: aws.String("tcp"),
			FromPort:   aws.Int64(testSSHPort),
			ToPort:     aws.Int64(testSSHPort),
		},
		{
			IpProtocol: aws.String("tcp"),
			FromPort:   aws.Int64(testBhojpurPort),
			ToPort:     aws.Int64(testBhojpurPort),
		},
	}

	perms, err := driver.configureSecurityGroupPermissions(group)

	assert.Nil(t, err)
	assert.Empty(t, perms)
}

func TestConfigureSecurityGroupPermissionsSkipReadOnly(t *testing.T) {
	driver := NewTestDriver()
	driver.SecurityGroupReadOnly = true
	perms, err := driver.configureSecurityGroupPermissions(securityGroupNoIpPermissions)

	assert.Nil(t, err)
	assert.Len(t, perms, 0)
}

func TestConfigureSecurityGroupPermissionsOpenPorts(t *testing.T) {
	driver := NewTestDriver()
	driver.OpenPorts = []string{"8888/tcp", "8080/udp", "9090"}
	perms, err := driver.configureSecurityGroupPermissions(securityGroupNoIpPermissions)

	assert.NoError(t, err)
	assert.Len(t, perms, 5)
	assert.Equal(t, aws.Int64(int64(8888)), perms[2].ToPort)
	assert.Equal(t, aws.String("tcp"), perms[2].IpProtocol)
	assert.Equal(t, aws.Int64(int64(8080)), perms[3].ToPort)
	assert.Equal(t, aws.String("udp"), perms[3].IpProtocol)
	assert.Equal(t, aws.Int64(int64(9090)), perms[4].ToPort)
	assert.Equal(t, aws.String("tcp"), perms[4].IpProtocol)
}

func TestConfigureSecurityGroupPermissionsOpenPortsSkipExisting(t *testing.T) {
	driver := NewTestDriver()
	group := securityGroup
	group.IpPermissions = []*ec2.IpPermission{
		{
			IpProtocol: aws.String("tcp"),
			FromPort:   aws.Int64(8888),
			ToPort:     aws.Int64(testSSHPort),
		},
		{
			IpProtocol: aws.String("tcp"),
			FromPort:   aws.Int64(8080),
			ToPort:     aws.Int64(testSSHPort),
		},
	}
	driver.OpenPorts = []string{"8888/tcp", "8080/udp", "8080"}
	perms, err := driver.configureSecurityGroupPermissions(group)
	assert.NoError(t, err)
	assert.Len(t, perms, 3)
	assert.Equal(t, aws.Int64(int64(8080)), perms[2].ToPort)
	assert.Equal(t, aws.String("udp"), perms[2].IpProtocol)
}

func TestConfigureSecurityGroupPermissionsInvalidOpenPorts(t *testing.T) {
	driver := NewTestDriver()
	driver.OpenPorts = []string{"2222/tcp", "abc1"}
	perms, err := driver.configureSecurityGroupPermissions(securityGroupNoIpPermissions)

	assert.Error(t, err)
	assert.Nil(t, perms)
}

func TestValidateAwsRegionValid(t *testing.T) {
	regions := []string{"eu-west-1", "eu-central-1"}

	for _, region := range regions {
		validatedRegion, err := validateAwsRegion(region)

		assert.NoError(t, err)
		assert.Equal(t, region, validatedRegion)
	}
}

func TestValidateAwsRegionInvalid(t *testing.T) {
	regions := []string{"eu-central-2"}

	for _, region := range regions {
		_, err := validateAwsRegion(region)

		assert.EqualError(t, err, "Invalid region specified")
	}
}

func TestFindDefaultVPC(t *testing.T) {
	driver := NewDriver("machineFoo", "path")
	driver.clientFactory = func() Ec2Client {
		return &fakeEC2WithLogin{}
	}

	vpc, err := driver.getDefaultVPCId()

	assert.Equal(t, "vpc-9999", vpc)
	assert.NoError(t, err)
}

func TestDefaultVPCIsMissing(t *testing.T) {
	driver := NewDriver("machineFoo", "path")
	driver.clientFactory = func() Ec2Client {
		return &fakeEC2WithDescribe{
			output: &ec2.DescribeAccountAttributesOutput{
				AccountAttributes: []*ec2.AccountAttribute{},
			},
		}
	}

	vpc, err := driver.getDefaultVPCId()

	assert.EqualError(t, err, "No default-vpc attribute")
	assert.Empty(t, vpc)
}

func TestGetRegionZoneForDefaultEndpoint(t *testing.T) {
	driver := NewCustomTestDriver(&fakeEC2WithLogin{})
	driver.awsCredentialsFactory = NewValidAwsCredentials
	options := &ctest.FakeFlagger{
		Data: map[string]interface{}{
			"name":             "test",
			"amazonec2-region": "us-east-1",
			"amazonec2-zone":   "e",
		},
	}

	err := driver.SetConfigFromFlags(options)

	regionZone := driver.getRegionZone()

	assert.Equal(t, "us-east-1e", regionZone)
	assert.NoError(t, err)
}

func TestGetRegionZoneForCustomEndpoint(t *testing.T) {
	driver := NewCustomTestDriver(&fakeEC2WithLogin{})
	driver.awsCredentialsFactory = NewValidAwsCredentials
	options := &ctest.FakeFlagger{
		Data: map[string]interface{}{
			"name":               "test",
			"amazonec2-endpoint": "https://someurl",
			"amazonec2-region":   "custom-endpoint",
			"amazonec2-zone":     "custom-zone",
		},
	}

	err := driver.SetConfigFromFlags(options)

	regionZone := driver.getRegionZone()

	assert.Equal(t, "custom-zone", regionZone)
	assert.NoError(t, err)
}

func TestDescribeAccountAttributeFails(t *testing.T) {
	driver := NewDriver("machineFoo", "path")
	driver.clientFactory = func() Ec2Client {
		return &fakeEC2WithDescribe{
			err: errors.New("Not Found"),
		}
	}

	vpc, err := driver.getDefaultVPCId()

	assert.EqualError(t, err, "Not Found")
	assert.Empty(t, vpc)
}

func TestAwsCredentialsAreRequired(t *testing.T) {
	driver := NewTestDriver()
	driver.awsCredentialsFactory = NewErrorAwsCredentials

	options := &ctest.FakeFlagger{
		Data: map[string]interface{}{
			"name":             "test",
			"amazonec2-region": "us-east-1",
			"amazonec2-zone":   "e",
		},
	}

	err := driver.SetConfigFromFlags(options)
	assert.Equal(t, err, errorMissingCredentials)
}

func TestValidAwsCredentialsAreAccepted(t *testing.T) {
	driver := NewCustomTestDriver(&fakeEC2WithLogin{})
	driver.awsCredentialsFactory = NewValidAwsCredentials
	options := &ctest.FakeFlagger{
		Data: map[string]interface{}{
			"name":             "test",
			"amazonec2-region": "us-east-1",
			"amazonec2-zone":   "e",
		},
	}

	err := driver.SetConfigFromFlags(options)
	assert.NoError(t, err)
}

func TestEndpointIsMandatoryWhenSSLDisabled(t *testing.T) {
	driver := NewTestDriver()
	driver.awsCredentialsFactory = NewValidAwsCredentials
	options := &ctest.FakeFlagger{
		Data: map[string]interface{}{
			"name":                         "test",
			"amazonec2-access-key":         "foobar",
			"amazonec2-region":             "us-east-1",
			"amazonec2-zone":               "e",
			"amazonec2-insecure-transport": true,
		},
	}

	err := driver.SetConfigFromFlags(options)

	assert.Equal(t, err, errorDisableSSLWithoutCustomEndpoint)
}

var values = []string{
	"bob",
	"jake",
	"jill",
}

var pointerSliceTests = []struct {
	input    []string
	expected []*string
}{
	{[]string{}, []*string{}},
	{[]string{values[1]}, []*string{&values[1]}},
	{[]string{values[0], values[2], values[2]}, []*string{&values[0], &values[2], &values[2]}},
}

func TestMakePointerSlice(t *testing.T) {
	for _, tt := range pointerSliceTests {
		actual := makePointerSlice(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}

var securityGroupNameTests = []struct {
	groupName  string
	groupNames []string
	expected   []string
}{
	{groupName: "bob", expected: []string{"bob"}},
	{groupNames: []string{"bill"}, expected: []string{"bill"}},
	{groupName: "bob", groupNames: []string{"bill"}, expected: []string{"bob", "bill"}},
}

func TestMergeSecurityGroupName(t *testing.T) {
	for _, tt := range securityGroupNameTests {
		d := Driver{SecurityGroupName: tt.groupName, SecurityGroupNames: tt.groupNames}
		assert.Equal(t, tt.expected, d.securityGroupNames())
	}
}

var securityGroupIdTests = []struct {
	groupId  string
	groupIds []string
	expected []string
}{
	{groupId: "id", expected: []string{"id"}},
	{groupIds: []string{"id"}, expected: []string{"id"}},
	{groupId: "id1", groupIds: []string{"id2"}, expected: []string{"id1", "id2"}},
}

func TestMergeSecurityGroupId(t *testing.T) {
	for _, tt := range securityGroupIdTests {
		d := Driver{SecurityGroupId: tt.groupId, SecurityGroupIds: tt.groupIds}
		assert.Equal(t, tt.expected, d.securityGroupIds())
	}
}

func matchGroupLookup(expected []string) interface{} {
	return func(input *ec2.DescribeSecurityGroupsInput) bool {
		actual := []string{}
		for _, filter := range input.Filters {
			if *filter.Name == "group-name" {
				for _, groupName := range filter.Values {
					actual = append(actual, *groupName)
				}
			}
		}
		return reflect.DeepEqual(expected, actual)
	}
}

func ipPermission(port int64) *ec2.IpPermission {
	return &ec2.IpPermission{
		FromPort:   aws.Int64(port),
		ToPort:     aws.Int64(port),
		IpProtocol: aws.String("tcp"),
		IpRanges:   []*ec2.IpRange{{CidrIp: aws.String(ipRange)}},
	}
}

func TestConfigureSecurityGroupsEmpty(t *testing.T) {
	recorder := fakeEC2SecurityGroupTestRecorder{}

	driver := NewCustomTestDriver(&recorder)
	err := driver.configureSecurityGroups([]string{})

	assert.Nil(t, err)
	recorder.AssertExpectations(t)
}

func TestConfigureSecurityGroupsMixed(t *testing.T) {
	groups := []string{"existingGroup", "newGroup"}
	recorder := fakeEC2SecurityGroupTestRecorder{}

	// First, a check is made for which groups already exist.
	initialLookupResult := ec2.DescribeSecurityGroupsOutput{SecurityGroups: []*ec2.SecurityGroup{
		{
			GroupName:     aws.String("existingGroup"),
			GroupId:       aws.String("existingGroupId"),
			IpPermissions: []*ec2.IpPermission{ipPermission(testSSHPort)},
		},
	}}
	recorder.On("DescribeSecurityGroups", mock.MatchedBy(matchGroupLookup(groups))).Return(
		&initialLookupResult, nil)

	// An ingress permission is added to the existing group.
	recorder.On("AuthorizeSecurityGroupIngress", &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:       aws.String("existingGroupId"),
		IpPermissions: []*ec2.IpPermission{ipPermission(testBhojpurPort)},
	}).Return(
		&ec2.AuthorizeSecurityGroupIngressOutput{}, nil)

	// The new security group is created.
	recorder.On("CreateSecurityGroup", &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String("newGroup"),
		Description: aws.String("Bhojpur Nodes"),
		VpcId:       aws.String(""),
	}).Return(
		&ec2.CreateSecurityGroupOutput{GroupId: aws.String("newGroupId")}, nil)

	// Ensuring the new security group exists.
	postCreateLookupResult := ec2.DescribeSecurityGroupsOutput{SecurityGroups: []*ec2.SecurityGroup{
		{
			GroupName: aws.String("newGroup"),
			GroupId:   aws.String("newGroupId"),
		},
	}}
	recorder.On("DescribeSecurityGroups",
		&ec2.DescribeSecurityGroupsInput{GroupIds: []*string{aws.String("newGroupId")}}).Return(
		&postCreateLookupResult, nil)

	// Permissions are added to the new security group.
	recorder.On("AuthorizeSecurityGroupIngress", &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:       aws.String("newGroupId"),
		IpPermissions: []*ec2.IpPermission{ipPermission(testSSHPort), ipPermission(testBhojpurPort)},
	}).Return(
		&ec2.AuthorizeSecurityGroupIngressOutput{}, nil)

	recorder.On("CreateTags", &ec2.CreateTagsInput{
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(machineTag),
				Value: aws.String(version.Version),
			},
		},
		Resources: []*string{aws.String("newGroupId")},
	}).Return(&ec2.CreateTagsOutput{}, nil)

	driver := NewCustomTestDriver(&recorder)
	err := driver.configureSecurityGroups(groups)

	assert.Nil(t, err)
	recorder.AssertExpectations(t)
}

func TestConfigureSecurityGroupsErrLookupExist(t *testing.T) {
	groups := []string{"group"}
	recorder := fakeEC2SecurityGroupTestRecorder{}

	lookupExistErr := errors.New("lookup failed")
	recorder.On("DescribeSecurityGroups", mock.MatchedBy(matchGroupLookup(groups))).Return(
		nil, lookupExistErr)

	driver := NewCustomTestDriver(&recorder)
	err := driver.configureSecurityGroups(groups)

	assert.Exactly(t, lookupExistErr, err)
	recorder.AssertExpectations(t)
}

func TestBase64UserDataIsEmptyIfNoFileProvided(t *testing.T) {
	driver := NewTestDriver()

	userdata, err := driver.Base64UserData()

	assert.NoError(t, err)
	assert.Empty(t, userdata)
}

func TestBase64UserDataGeneratesErrorIfFileNotFound(t *testing.T) {
	dir, err := ioutil.TempDir("", "awsuserdata")
	assert.NoError(t, err, "Unable to create temporary directory.")

	defer os.RemoveAll(dir)
	userdata_path := filepath.Join(dir, "does-not-exist.yml")

	driver := NewTestDriver()
	driver.UserDataFile = userdata_path

	_, ud_err := driver.Base64UserData()
	assert.Equal(t, ud_err, errorReadingUserData)
}

func TestBase64UserDataIsCorrectWhenFileProvided(t *testing.T) {
	dir, err := ioutil.TempDir("", "awsuserdata")
	assert.NoError(t, err, "Unable to create temporary directory.")

	defer os.RemoveAll(dir)

	userdata_path := filepath.Join(dir, "test-userdata.yml")

	content := []byte("#cloud-config\nhostname: userdata-test\nfqdn: userdata-test.amazonec2.driver\n")
	contentBase64 := "I2Nsb3VkLWNvbmZpZwpob3N0bmFtZTogdXNlcmRhdGEtdGVzdApmcWRuOiB1c2VyZGF0YS10ZXN0LmFtYXpvbmVjMi5kcml2ZXIK"

	err = ioutil.WriteFile(userdata_path, content, 0666)
	assert.NoError(t, err, "Unable to create temporary userdata file.")

	driver := NewTestDriver()
	driver.UserDataFile = userdata_path

	userdata, ud_err := driver.Base64UserData()

	assert.NoError(t, ud_err)
	assert.Equal(t, contentBase64, userdata)
}

func TestDefaultAMI(t *testing.T) {
	driver := NewCustomTestDriver(&fakeEC2WithLogin{})

	err := driver.checkAMI()

	assert.Equal(t, "/dev/sda1", driver.DeviceName)
	assert.NoError(t, err)
}

func TestRootDeviceName(t *testing.T) {
	driver := NewCustomTestDriver(&fakeEC2WithLogin{})
	driver.AMI = "ami-0eeb1ef502d7b850d" // Fedora CoreOS image

	err := driver.checkAMI()

	assert.Equal(t, "/dev/xvda", driver.DeviceName)
	assert.NoError(t, err)
}

func TestInvalidAMI(t *testing.T) {
	driver := NewCustomTestDriver(&fakeEC2WithLogin{})
	driver.AMI = "ami-000" // Invalid AMI

	err := driver.checkAMI()

	assert.Error(t, err)
}
