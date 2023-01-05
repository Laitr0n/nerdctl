/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package types

type ComposeBuildCommandOptions struct {
	// BuildArgs is a list of build-time variables for services.
	BuildArgs []string
	// NoCache is a flag to disable cache when building the image.
	NoCache bool
	// Progress is a flag to show the type of progress output (auto, plain, tty)
	Progress string
}

type ComposeConfigCommandOptions struct {
	// Quiet is a flag to only validate the configuration, without printing.
	Quiet bool
	// Services is flag to print the service names, one per line.
	Services bool
	// Volumes is a flag to print the volume names, one per line.
	Volumes bool
	// Hash is a flag to print the service configuration hash, one per line.
	Hash string
}
