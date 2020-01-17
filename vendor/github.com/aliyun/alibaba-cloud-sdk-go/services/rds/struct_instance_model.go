package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// InstanceModel is a nested struct in rds response
type InstanceModel struct {
	DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
	Region           string `json:"Region" xml:"Region"`
	ZoneId           string `json:"ZoneId" xml:"ZoneId"`
	Engine           string `json:"Engine" xml:"Engine"`
	PayType          string `json:"PayType" xml:"PayType"`
	CreatedTime      string `json:"CreatedTime" xml:"CreatedTime"`
	ExpireTime       string `json:"ExpireTime" xml:"ExpireTime"`
	LockMode         string `json:"LockMode" xml:"LockMode"`
	DBInstanceStatus string `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
}