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

type BinLogFile struct {
	FileSize             int    `json:"FileSize" xml:"FileSize"`
	LogBeginTime         string `json:"LogBeginTime" xml:"LogBeginTime"`
	LogEndTime           string `json:"LogEndTime" xml:"LogEndTime"`
	DownloadLink         string `json:"DownloadLink" xml:"DownloadLink"`
	IntranetDownloadLink string `json:"IntranetDownloadLink" xml:"IntranetDownloadLink"`
	LinkExpiredTime      string `json:"LinkExpiredTime" xml:"LinkExpiredTime"`
	Checksum             string `json:"Checksum" xml:"Checksum"`
	HostInstanceID       string `json:"HostInstanceID" xml:"HostInstanceID"`
}
