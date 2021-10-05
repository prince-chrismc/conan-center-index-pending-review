package pending_review

import (
	"encoding/json"
	"testing"

	"github.com/google/go-github/v39/github"
	"github.com/stretchr/testify/assert"
)

func parsePrJson(t *testing.T, str string) []*github.CommitFile {
	var files []*github.CommitFile

	if err := json.Unmarshal([]byte(str), &files); err != nil {
		t.Fatal()
	}

	return files
}

func TestOnlyBumpFilesChangedNotTwo(t *testing.T) {
	oneFile := parsePrJson(t, `[
		{
		  "sha": "5cbce65d888e970205160de1ea33cb3dae4b948b",
		  "filename": "recipes/b2/portable/conandata.yml",
		  "status": "modified",
		  "additions": 3,
		  "deletions": 0,
		  "changes": 3,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/7558ff23fa9eabd5ae08e90b89abc125f4a557e4/recipes/b2/portable/conandata.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/7558ff23fa9eabd5ae08e90b89abc125f4a557e4/recipes/b2/portable/conandata.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes/b2/portable/conandata.yml?ref=7558ff23fa9eabd5ae08e90b89abc125f4a557e4",
		  "patch": "@@ -17,3 +17,6 @@ sources:\n   \"4.6.0\":\n     url: \"https://github.com/bfgroup/b2/archive/4.6.0.tar.gz\"\n     sha256: \"3a308e0f79a039d8a9495b375f3292f5163000c19caa79c5687e4cb5b1938b49\"\n+  \"4.6.1\":\n+    url: \"https://github.com/bfgroup/b2/archive/4.6.1.tar.gz\"\n+    sha256: \"a3f3323eaeb2c27d7a3ca86842665c6c3bc3d93cc626ba362ae6d0c5a7bfbe2c\""
		}
	  ]`)

	assert.Equal(t, false, onlyBumpFilesChanged(oneFile))
}

func TestOnlyBumpFilesChangedWrongFiles(t *testing.T) {
	filesCMakeRecipe := parsePrJson(t, `[
		{
		  "sha": "9c7b7521252773b5e1880c9c7f4cbc0f2196ad42",
		  "filename": "recipes/cpu_features/all/CMakeLists.txt",
		  "status": "modified",
		  "additions": 1,
		  "deletions": 1,
		  "changes": 2,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/480a685625a1726f9a7685b8f0f96b12871eb346/recipes/cpu_features/all/CMakeLists.txt",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/480a685625a1726f9a7685b8f0f96b12871eb346/recipes/cpu_features/all/CMakeLists.txt",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes/cpu_features/all/CMakeLists.txt?ref=480a685625a1726f9a7685b8f0f96b12871eb346",
		  "patch": "@@ -3,7 +3,7 @@ project(cmake_wrapper)\n \n set(CMAKE_WINDOWS_EXPORT_ALL_SYMBOLS ON)\n \n-include(\"${CMAKE_BINARY_DIR}/conanbuildinfo.cmake\")\n+include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)\n conan_basic_setup()\n \n add_subdirectory(\"source_subfolder\")"
		},
		{
		  "sha": "b7346aaae30a5697f86f55f707fe1f5b81166afa",
		  "filename": "recipes/cpu_features/all/conanfile.py",
		  "status": "modified",
		  "additions": 7,
		  "deletions": 12,
		  "changes": 19,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/480a685625a1726f9a7685b8f0f96b12871eb346/recipes/cpu_features/all/conanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/480a685625a1726f9a7685b8f0f96b12871eb346/recipes/cpu_features/all/conanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes/cpu_features/all/conanfile.py?ref=480a685625a1726f9a7685b8f0f96b12871eb346",
		  "patch": "@@ -2,6 +2,8 @@\n from conans import ConanFile, CMake, tools\n from conans.errors import ConanInvalidConfiguration\n \n+required_conan_version = \">=1.33.0\"\n+\n \n class CpuFeaturesConan(ConanFile):\n     name = \"cpu_features\"\n@@ -11,10 +13,8 @@ class CpuFeaturesConan(ConanFile):\n     description = \"A cross platform C99 library to get cpu features at runtime.\"\n     topics = (\"conan\", \"cpu\", \"features\", \"cpuid\")\n     settings = \"os\", \"arch\", \"compiler\", \"build_type\"\n-    options = {\"shared\": [True, False],\n-               \"fPIC\": [True, False]}\n-    default_options = {\"shared\": False,\n-                       \"fPIC\": True}\n+    options = {\"shared\": [True, False], \"fPIC\": [True, False]}\n+    default_options = {\"shared\": False, \"fPIC\": True}\n     exports_sources = [\"CMakeLists.txt\"]\n     generators = \"cmake\",\n     _cmake = None\n@@ -23,10 +23,6 @@ class CpuFeaturesConan(ConanFile):\n     def _source_subfolder(self):\n         return \"source_subfolder\"\n \n-    @property\n-    def _build_subfolder(self):\n-        return \"build_subfolder\"\n-\n     def source(self):\n         tools.get(**self.conan_data[\"sources\"][self.version], strip_root=True, destination=self._source_subfolder)\n \n@@ -47,10 +43,9 @@ def config_options(self):\n     def _configure_cmake(self):\n         if self._cmake:\n             return self._cmake\n-        cmake = CMake(self)\n-        cmake.definitions[\"BUILD_PIC\"] = self.options.get_safe(\"fPIC\", True)\n-        cmake.configure()\n-        self._cmake = cmake\n+        self._cmake = CMake(self)\n+        self._cmake.definitions[\"BUILD_PIC\"] = self.options.get_safe(\"fPIC\", True)\n+        self._cmake.configure() # Does not support out of source builds\n         return self._cmake\n \n     def build(self):"
		}
	  ]`)

	assert.Equal(t, false, onlyBumpFilesChanged(filesCMakeRecipe))
}

func TestOnlyBumpFilesChanged(t *testing.T) {
	filesConfigData := parsePrJson(t, `[
		{
		  "sha": "d4cfe969ef2e75f2f66cfbd4e41b61ee50962d54",
		  "filename": "recipes/b2/config.yml",
		  "status": "modified",
		  "additions": 2,
		  "deletions": 0,
		  "changes": 2,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/7558ff23fa9eabd5ae08e90b89abc125f4a557e4/recipes/b2/config.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/7558ff23fa9eabd5ae08e90b89abc125f4a557e4/recipes/b2/config.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes/b2/config.yml?ref=7558ff23fa9eabd5ae08e90b89abc125f4a557e4",
		  "patch": "@@ -19,3 +19,5 @@ versions:\n     folder: portable\n   \"4.6.0\":\n     folder: portable\n+  \"4.6.1\":\n+    folder: portable"
		},
		{
		  "sha": "5cbce65d888e970205160de1ea33cb3dae4b948b",
		  "filename": "recipes/b2/portable/conandata.yml",
		  "status": "modified",
		  "additions": 3,
		  "deletions": 0,
		  "changes": 3,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/7558ff23fa9eabd5ae08e90b89abc125f4a557e4/recipes/b2/portable/conandata.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/7558ff23fa9eabd5ae08e90b89abc125f4a557e4/recipes/b2/portable/conandata.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes/b2/portable/conandata.yml?ref=7558ff23fa9eabd5ae08e90b89abc125f4a557e4",
		  "patch": "@@ -17,3 +17,6 @@ sources:\n   \"4.6.0\":\n     url: \"https://github.com/bfgroup/b2/archive/4.6.0.tar.gz\"\n     sha256: \"3a308e0f79a039d8a9495b375f3292f5163000c19caa79c5687e4cb5b1938b49\"\n+  \"4.6.1\":\n+    url: \"https://github.com/bfgroup/b2/archive/4.6.1.tar.gz\"\n+    sha256: \"a3f3323eaeb2c27d7a3ca86842665c6c3bc3d93cc626ba362ae6d0c5a7bfbe2c\""
		}
	  ]`)

	assert.Equal(t, true, onlyBumpFilesChanged(filesConfigData))
}

func TestOnlyBumpFilesChangedWorngOrder(t *testing.T) {
	filesDataConfig := parsePrJson(t, `[
		{
		  "sha": "b7081bc87d34dcf2cebac5e897e4869fc92d71a9",
		  "filename": "recipes/djinni-generator/all/conandata.yml",
		  "status": "modified",
		  "additions": 3,
		  "deletions": 0,
		  "changes": 3,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/83f2e29822377fd89d7cb4d3a2be50ceab2c269a/recipes/djinni-generator/all/conandata.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/83f2e29822377fd89d7cb4d3a2be50ceab2c269a/recipes/djinni-generator/all/conandata.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes/djinni-generator/all/conandata.yml?ref=83f2e29822377fd89d7cb4d3a2be50ceab2c269a",
		  "patch": "@@ -14,3 +14,6 @@ sources:\n   \"1.0.0\":\n     url: https://github.com/cross-language-cpp/djinni-generator/releases/download/v1.0.0/djinni\n     sha256: \"a5dc94cd5175f228eb17e93e0f4cec93c18615758b139707a05f20dc7001f55f\"\n+  \"1.1.0\":\n+    url: https://github.com/cross-language-cpp/djinni-generator/releases/download/v1.1.0/djinni\n+    sha256: \"4efd4f68cf913af7c9dd7dd975a8aa2d2a90e8efd9d3b556979ff577923e0d44\""
		},
		{
		  "sha": "8af4fbf3c504e083e84fa832469a252b1bcc74fa",
		  "filename": "recipes/djinni-generator/config.yml",
		  "status": "modified",
		  "additions": 2,
		  "deletions": 0,
		  "changes": 2,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/83f2e29822377fd89d7cb4d3a2be50ceab2c269a/recipes/djinni-generator/config.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/83f2e29822377fd89d7cb4d3a2be50ceab2c269a/recipes/djinni-generator/config.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes/djinni-generator/config.yml?ref=83f2e29822377fd89d7cb4d3a2be50ceab2c269a",
		  "patch": "@@ -9,3 +9,5 @@ versions:\n     folder: \"all\"\n   \"1.0.0\":\n     folder: \"all\"\n+  \"1.1.0\":\n+    folder: \"all\""
		}
	  ]`)

	assert.Equal(t, true, onlyBumpFilesChanged(filesDataConfig))
}
