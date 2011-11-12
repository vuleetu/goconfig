// Copyright 2010  The "goconfig" Authors
//
// Use of this source code is governed by the BSD 2-Clause License
// that can be found in the LICENSE file.
//
// This software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
// OR CONDITIONS OF ANY KIND, either express or implied. See the License
// for more details.

package config

type sectionError string

func (self sectionError) Error() string {
	return "section not found: " + string(self)
}

type optionError string

func (self optionError) Error() string {
	return "option not found: " + string(self)
}
