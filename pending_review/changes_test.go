package pending_review

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parsePullRequestFilesJSON(t *testing.T, str string) []*CommitFile {
	var files []*CommitFile

	if err := json.Unmarshal([]byte(str), &files); err != nil {
		t.Fatal("Unable to parse JSON")
	}

	return files
}

func TestProcessChangedFilesTinyEdit(t *testing.T) {
	// https://github.com/conan-io/conan-center-index/pull/15470/files
	files := parsePullRequestFilesJSON(t, `[
    {
      "sha": "6761b154e37b6cdeac8a4117c420b3a3ab1380cc",
      "filename": "recipes/jfalcou-eve/all/conanfile.py",
      "status": "modified",
      "additions": 1,
      "deletions": 3,
      "changes": 4,
      "blob_url": "https://github.com/conan-io/conan-center-index/blob/8a13d1497e84d700ffe2a8560a0d7a8c9e2a64ee/recipes%2Fjfalcou-eve%2Fall%2Fconanfile.py",
      "raw_url": "https://github.com/conan-io/conan-center-index/raw/8a13d1497e84d700ffe2a8560a0d7a8c9e2a64ee/recipes%2Fjfalcou-eve%2Fall%2Fconanfile.py",
      "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fjfalcou-eve%2Fall%2Fconanfile.py?ref=8a13d1497e84d700ffe2a8560a0d7a8c9e2a64ee",
      "patch": "@@ -35,7 +35,7 @@ def _compilers_minimum_version(self):\n                 \"Visual Studio\": \"16.9\",\n                 \"msvc\": \"1928\",\n                 \"clang\": \"13\",\n-                \"apple-clang\": \"13\"}\n+                \"apple-clang\": \"14\"}\n \n     def configure(self):\n         version = Version(self.version.strip(\"v\"))\n@@ -53,8 +53,6 @@ def validate(self):\n             check_min_cppstd(self, self._min_cppstd)\n         if is_msvc(self):\n             raise ConanInvalidConfiguration(\"EVE does not support MSVC yet (https://github.com/jfalcou/eve/issues/1022).\")\n-        if self.settings.compiler == \"apple-clang\":\n-            raise ConanInvalidConfiguration(\"EVE does not support apple Clang due to an incomplete libcpp.\")\n \n         def lazy_lt_semver(v1, v2):\n             lv1 = [int(v) for v in v1.split(\".\")]"
    }
  ]`)

	obtained, err := processChangedFiles(files)
	assert.Equal(t, err, nil)
	assert.Equal(t, &change{Recipe: "jfalcou-eve", Change: EDIT, Weight: TINY}, obtained)
}

func TestProcessChangedFilesRegularEdit(t *testing.T) {
	// https://github.com/conan-io/conan-center-index/pull/15372/files
	files := parsePullRequestFilesJSON(t, `[
		{
		  "sha": "b0d1e5271b8503adea65c7477489ad7b77a1b043",
		  "filename": "recipes/flac/all/conandata.yml",
		  "status": "modified",
		  "additions": 10,
		  "deletions": 1,
		  "changes": 11,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/955ef84abb67216777c78a1c73ba59e8370b0937/recipes%2Fflac%2Fall%2Fconandata.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/955ef84abb67216777c78a1c73ba59e8370b0937/recipes%2Fflac%2Fall%2Fconandata.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fflac%2Fall%2Fconandata.yml?ref=955ef84abb67216777c78a1c73ba59e8370b0937",
		  "patch": "@@ -1,7 +1,16 @@\n sources:\n+  \"1.4.2\":\n+    url: \"https://github.com/xiph/flac/releases/download/1.4.2/flac-1.4.2.tar.xz\"\n+    sha256: \"e322d58a1f48d23d9dd38f432672865f6f79e73a6f9cc5a5f57fcaa83eb5a8e4\"\n   \"1.3.3\":\n     url: \"https://github.com/xiph/flac/archive/1.3.3.tar.gz\"\n     sha256: \"668cdeab898a7dd43cf84739f7e1f3ed6b35ece2ef9968a5c7079fe9adfe1689\"\n patches:\n+  \"1.4.2\":\n+    - patch_file: \"patches/fix-cmake-1.4.2.patch\"\n+      patch_description: \"Adapts find_package commands and install destination paths in CMakeLists.txt files.\"\n+      patch_type: \"conan\"\n   \"1.3.3\":\n-    - patch_file: \"patches/fix-cmake.patch\"\n+    - patch_file: \"patches/fix-cmake-1.3.3.patch\"\n+      patch_description: \"Various adaptations in CMakeLists.txt files to improve compatibility with Conan.\"\n+      patch_type: \"conan\""
		},
		{
		  "sha": "3a9e705b2a4f7c0ff46147bcac198539b61fe62f",
		  "filename": "recipes/flac/all/conanfile.py",
		  "status": "modified",
		  "additions": 10,
		  "deletions": 7,
		  "changes": 17,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/955ef84abb67216777c78a1c73ba59e8370b0937/recipes%2Fflac%2Fall%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/955ef84abb67216777c78a1c73ba59e8370b0937/recipes%2Fflac%2Fall%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fflac%2Fall%2Fconanfile.py?ref=955ef84abb67216777c78a1c73ba59e8370b0937",
		  "patch": "@@ -1,10 +1,11 @@\n from conan import ConanFile\n from conan.tools.cmake import CMake, CMakeDeps, CMakeToolchain, cmake_layout\n from conan.tools.env import VirtualBuildEnv\n-from conan.tools.files import apply_conandata_patches, copy, get, rmdir\n+from conan.tools.files import apply_conandata_patches, export_conandata_patches, copy, get, rmdir\n+from conan.tools.scm import Version\n import os\n \n-required_conan_version = \">=1.47.0\"\n+required_conan_version = \">=1.53.0\"\n \n \n class FlacConan(ConanFile):\n@@ -26,22 +27,21 @@ class FlacConan(ConanFile):\n     }\n \n     def export_sources(self):\n-        for p in self.conan_data.get(\"patches\", {}).get(self.version, []):\n-            copy(self, p[\"patch_file\"], self.recipe_folder, self.export_sources_folder)\n+        export_conandata_patches(self)\n \n     def config_options(self):\n         if self.settings.os == \"Windows\":\n             del self.options.fPIC\n \n     def configure(self):\n         if self.options.shared:\n-            del self.options.fPIC\n+            self.options.rm_safe(\"fPIC\")\n \n     def requirements(self):\n         self.requires(\"ogg/1.3.5\")\n \n     def build_requirements(self):\n-        if self.settings.arch in [\"x86\", \"x86_64\"]:\n+        if Version(self.version) < \"1.4.2\" and self.settings.arch in [\"x86\", \"x86_64\"]:\n             self.tool_requires(\"nasm/2.15.05\")\n \n     def layout(self):\n@@ -56,6 +56,8 @@ def generate(self):\n         tc.variables[\"BUILD_EXAMPLES\"] = False\n         tc.variables[\"BUILD_DOCS\"] = False\n         tc.variables[\"BUILD_TESTING\"] = False\n+        # Honor BUILD_SHARED_LIBS from conan_toolchain (see https://github.com/conan-io/conan/issues/11840)\n+        tc.cache_variables[\"CMAKE_POLICY_DEFAULT_CMP0077\"] = \"NEW\"\n         tc.generate()\n         cd = CMakeDeps(self)\n         cd.generate()\n@@ -79,6 +81,8 @@ def package(self):\n         copy(self, \"*.h\", src=os.path.join(self.source_folder, \"include\", \"share\", \"grabbag\"),\n                           dst=os.path.join(self.package_folder, \"include\", \"share\", \"grabbag\"), keep_path=False)\n         rmdir(self, os.path.join(self.package_folder, \"share\"))\n+        rmdir(self, os.path.join(self.package_folder, \"lib\", \"cmake\"))\n+        rmdir(self, os.path.join(self.package_folder, \"lib\", \"pkgconfig\"))\n \n     def package_info(self):\n         self.cpp_info.set_property(\"cmake_file_name\", \"flac\")\n@@ -101,7 +105,6 @@ def package_info(self):\n         self.output.info(\"Appending PATH environment variable: {}\".format(bin_path))\n         self.env_info.PATH.append(bin_path)\n \n-        # TODO: to remove in conan v2 once cmake_find_package_* generators removed\n         self.cpp_info.filenames[\"cmake_find_package\"] = \"flac\"\n         self.cpp_info.filenames[\"cmake_find_package_multi\"] = \"flac\"\n         self.cpp_info.names[\"cmake_find_package\"] = \"FLAC\""
		},
		{
		  "sha": "ec7db2f2d00aaaa31bcbebdc536d8faadb43c84b",
		  "filename": "recipes/flac/all/patches/fix-cmake-1.3.3.patch",
		  "status": "renamed",
		  "additions": 0,
		  "deletions": 0,
		  "changes": 0,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/955ef84abb67216777c78a1c73ba59e8370b0937/recipes%2Fflac%2Fall%2Fpatches%2Ffix-cmake-1.3.3.patch",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/955ef84abb67216777c78a1c73ba59e8370b0937/recipes%2Fflac%2Fall%2Fpatches%2Ffix-cmake-1.3.3.patch",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fflac%2Fall%2Fpatches%2Ffix-cmake-1.3.3.patch?ref=955ef84abb67216777c78a1c73ba59e8370b0937",
		  "previous_filename": "recipes/flac/all/patches/fix-cmake.patch"
		},
		{
		  "sha": "bd5a0ebdb6997e0c7b0a71f0bd93830abd2266a5",
		  "filename": "recipes/flac/all/patches/fix-cmake-1.4.2.patch",
		  "status": "added",
		  "additions": 38,
		  "deletions": 0,
		  "changes": 38,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/955ef84abb67216777c78a1c73ba59e8370b0937/recipes%2Fflac%2Fall%2Fpatches%2Ffix-cmake-1.4.2.patch",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/955ef84abb67216777c78a1c73ba59e8370b0937/recipes%2Fflac%2Fall%2Fpatches%2Ffix-cmake-1.4.2.patch",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fflac%2Fall%2Fpatches%2Ffix-cmake-1.4.2.patch?ref=955ef84abb67216777c78a1c73ba59e8370b0937",
		  "patch": "@@ -0,0 +1,38 @@\n+--- a/CMakeLists.txt\n++++ b/CMakeLists.txt\n+@@ -43,7 +43,7 @@ if(WITH_OGG)\n+         endif()\n+     else()\n+         if(NOT TARGET Ogg::ogg)\n+-            find_package(Ogg REQUIRED)\n++            find_package(Ogg REQUIRED CONFIG)\n+         else()\n+             set(OGG_FOUND 1 CACHE INTERNAL \"ogg has already been built\")\n+         endif()\n+--- a/src/flac/CMakeLists.txt\n++++ b/src/flac/CMakeLists.txt\n+@@ -21,4 +21,4 @@ target_link_libraries(flacapp\n+     utf8)\n+ \n+ install(TARGETS flacapp EXPORT targets\n+-    RUNTIME DESTINATION \"${CMAKE_INSTALL_BINDIR}\")\n++    DESTINATION \"${CMAKE_INSTALL_BINDIR}\")\n+--- a/src/metaflac/CMakeLists.txt\n++++ b/src/metaflac/CMakeLists.txt\n+@@ -14,4 +14,4 @@ add_executable(metaflac\n+ target_link_libraries(metaflac FLAC getopt utf8)\n+ \n+ install(TARGETS metaflac EXPORT targets\n+-    RUNTIME DESTINATION \"${CMAKE_INSTALL_BINDIR}\")\n++    DESTINATION \"${CMAKE_INSTALL_BINDIR}\")\n+--- a/src/share/getopt/CMakeLists.txt\n++++ b/src/share/getopt/CMakeLists.txt\n+@@ -1,8 +1,7 @@\n+ check_include_file(\"string.h\" HAVE_STRING_H)\n+ \n+ if(NOT WIN32)\n+-    find_package(Intl)\n+ endif()\n+ \n+ add_library(getopt STATIC getopt.c getopt1.c)\n+ "
		},
		{
		  "sha": "fac8ae5f33f51941a2c2fac89f8e8d2a4508a6e0",
		  "filename": "recipes/flac/config.yml",
		  "status": "modified",
		  "additions": 2,
		  "deletions": 0,
		  "changes": 2,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/955ef84abb67216777c78a1c73ba59e8370b0937/recipes%2Fflac%2Fconfig.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/955ef84abb67216777c78a1c73ba59e8370b0937/recipes%2Fflac%2Fconfig.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fflac%2Fconfig.yml?ref=955ef84abb67216777c78a1c73ba59e8370b0937",
		  "patch": "@@ -1,3 +1,5 @@\n versions:\n+  \"1.4.2\":\n+    folder: all\n   \"1.3.3\":\n     folder: all"
		}
	  ]`)

	obtained, err := processChangedFiles(files)
	assert.Equal(t, err, nil)
	assert.Equal(t, &change{Recipe: "flac", Change: EDIT, Weight: REGULAR}, obtained)
}

func TestProcessChangedFilesHeavyEdit(t *testing.T) {
	// https://github.com/conan-io/conan-center-index/pull/15350/files
	files := parsePullRequestFilesJSON(t, `[
		{
		  "sha": "b71c882d9d33fc382d6af137c1acc5f3b7fc91a7",
		  "filename": "recipes/libfdk_aac/all/CMakeLists.txt",
		  "status": "removed",
		  "additions": 0,
		  "deletions": 7,
		  "changes": 7,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/d6e2e52ebfa83281c5ad6f91e91cfaffcc73e1d5/recipes%2Flibfdk_aac%2Fall%2FCMakeLists.txt",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/d6e2e52ebfa83281c5ad6f91e91cfaffcc73e1d5/recipes%2Flibfdk_aac%2Fall%2FCMakeLists.txt",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Flibfdk_aac%2Fall%2FCMakeLists.txt?ref=d6e2e52ebfa83281c5ad6f91e91cfaffcc73e1d5",
		  "patch": "@@ -1,7 +0,0 @@\n-cmake_minimum_required(VERSION 3.1)\n-project(cmake_wrapper)\n-\n-include(conanbuildinfo.cmake)\n-conan_basic_setup(KEEP_RPATHS)\n-\n-add_subdirectory(source_subfolder)"
		},
		{
		  "sha": "5907c10a61c22c2dec1311ec8d0f8ba60bc423d7",
		  "filename": "recipes/libfdk_aac/all/conanfile.py",
		  "status": "modified",
		  "additions": 88,
		  "deletions": 103,
		  "changes": 191,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/893ee53357b1914124b535113be20162e47f52cc/recipes%2Flibfdk_aac%2Fall%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/893ee53357b1914124b535113be20162e47f52cc/recipes%2Flibfdk_aac%2Fall%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Flibfdk_aac%2Fall%2Fconanfile.py?ref=893ee53357b1914124b535113be20162e47f52cc",
		  "patch": "@@ -1,9 +1,15 @@\n-from conans import ConanFile, AutoToolsBuildEnvironment, CMake, VisualStudioBuildEnvironment, tools\n-import contextlib\n-import functools\n+from conan import ConanFile\n+from conan.tools.apple import fix_apple_shared_install_name\n+from conan.tools.cmake import CMake, CMakeToolchain, cmake_layout\n+from conan.tools.env import VirtualBuildEnv\n+from conan.tools.files import chdir, copy, get, rename, replace_in_file, rm, rmdir\n+from conan.tools.gnu import Autotools, AutotoolsToolchain\n+from conan.tools.layout import basic_layout\n+from conan.tools.microsoft import is_msvc, NMakeToolchain\n+from conan.tools.scm import Version\n import os\n \n-required_conan_version = \">=1.43.0\"\n+required_conan_version = \">=1.55.0\"\n \n \n class LibFDKAACConan(ConanFile):\n@@ -12,7 +18,7 @@ class LibFDKAACConan(ConanFile):\n     description = \"A standalone library of the Fraunhofer FDK AAC code from Android\"\n     license = \"https://github.com/mstorsjo/fdk-aac/blob/master/NOTICE\"\n     homepage = \"https://sourceforge.net/projects/opencore-amr/\"\n-    topics = (\"libfdk_aac\", \"multimedia\", \"audio\", \"fraunhofer\", \"aac\", \"decoder\", \"encoding\", \"decoding\")\n+    topics = (\"multimedia\", \"audio\", \"fraunhofer\", \"aac\", \"decoder\", \"encoding\", \"decoding\")\n \n     settings = \"os\", \"arch\", \"compiler\", \"build_type\"\n     options = {\n@@ -24,133 +30,113 @@ class LibFDKAACConan(ConanFile):\n         \"fPIC\": True,\n     }\n \n-    exports_sources = \"CMakeLists.txt\"\n-    generators = \"cmake\"\n-\n-    @property\n-    def _source_subfolder(self):\n-        return \"source_subfolder\"\n-\n-    @property\n-    def _is_msvc(self):\n-        return str(self.settings.compiler) in [\"Visual Studio\", \"msvc\"]\n-\n     @property\n     def _settings_build(self):\n         return getattr(self, \"settings_build\", self.settings)\n \n     @property\n     def _use_cmake(self):\n-        return tools.Version(self.version) >= \"2.0.2\"\n+        return Version(self.version) >= \"2.0.2\"\n \n     def config_options(self):\n         if self.settings.os == \"Windows\":\n             del self.options.fPIC\n \n     def configure(self):\n         if self.options.shared:\n-            del self.options.fPIC\n+            self.options.rm_safe(\"fPIC\")\n+\n+    def layout(self):\n+        if self._use_cmake:\n+            cmake_layout(self, src_folder=\"src\")\n+        else:\n+            basic_layout(self, src_folder=\"src\")\n \n     def build_requirements(self):\n-        if not self._use_cmake and not self._is_msvc:\n-            self.build_requires(\"libtool/2.4.6\")\n-            if self._settings_build.os == \"Windows\" and not tools.get_env(\"CONAN_BASH_PATH\"):\n-                self.build_requires(\"msys2/cci.latest\")\n+        if not self._use_cmake and not is_msvc(self):\n+            self.tool_requires(\"libtool/2.4.7\")\n+            if self._settings_build.os == \"Windows\":\n+                self.win_bash = True\n+                if not self.conf.get(\"tools.microsoft.bash:path\", check_type=str):\n+                    self.tool_requires(\"msys2/cci.latest\")\n \n     def source(self):\n-        tools.get(**self.conan_data[\"sources\"][self.version],\n-                  destination=self._source_subfolder, strip_root=True)\n-\n-    @functools.lru_cache(1)\n-    def _configure_cmake(self):\n-        cmake = CMake(self)\n-        cmake.definitions[\"BUILD_PROGRAMS\"] = False\n-        cmake.definitions[\"FDK_AAC_INSTALL_CMAKE_CONFIG_MODULE\"] = False\n-        cmake.definitions[\"FDK_AAC_INSTALL_PKGCONFIG_MODULE\"] = False\n-        cmake.configure()\n-        return cmake\n-\n-    @contextlib.contextmanager\n-    def _msvc_build_environment(self):\n-        with tools.chdir(self._source_subfolder):\n-            with tools.vcvars(self):\n-                with tools.environment_append(VisualStudioBuildEnvironment(self).vars):\n-                    yield\n-\n-    def _build_vs(self):\n-        with self._msvc_build_environment():\n-            # Rely on flags injected by conan\n-            tools.replace_in_file(\"Makefile.vc\",\n-                                  \"CFLAGS   = /nologo /W3 /Ox /MT\",\n-                                  \"CFLAGS   = /nologo\")\n-            tools.replace_in_file(\"Makefile.vc\",\n-                                  \"MKDIR_FLAGS = -p\",\n-                                  \"MKDIR_FLAGS =\")\n-            # Build either shared or static, and don't build utility (it always depends on static lib)\n-            tools.replace_in_file(\"Makefile.vc\", \"copy $(PROGS) $(bindir)\", \"\")\n-            tools.replace_in_file(\"Makefile.vc\", \"copy $(LIB_DEF) $(libdir)\", \"\")\n-            if self.options.shared:\n-                tools.replace_in_file(\"Makefile.vc\",\n-                                      \"all: $(LIB_DEF) $(STATIC_LIB) $(SHARED_LIB) $(IMP_LIB) $(PROGS)\",\n-                                      \"all: $(LIB_DEF) $(SHARED_LIB) $(IMP_LIB)\")\n-                tools.replace_in_file(\"Makefile.vc\", \"copy $(STATIC_LIB) $(libdir)\", \"\")\n-            else:\n-                tools.replace_in_file(\"Makefile.vc\",\n-                                      \"all: $(LIB_DEF) $(STATIC_LIB) $(SHARED_LIB) $(IMP_LIB) $(PROGS)\",\n-                                      \"all: $(STATIC_LIB)\")\n-                tools.replace_in_file(\"Makefile.vc\", \"copy $(IMP_LIB) $(libdir)\", \"\")\n-                tools.replace_in_file(\"Makefile.vc\", \"copy $(SHARED_LIB) $(bindir)\", \"\")\n-            self.run(\"nmake -f Makefile.vc\")\n-\n-    def _build_autotools(self):\n-        with tools.chdir(self._source_subfolder):\n-            self.run(\"{} -fiv\".format(tools.get_env(\"AUTORECONF\")), win_bash=tools.os_info.is_windows)\n-            # relocatable shared lib on macOS\n-            tools.replace_in_file(\"configure\", \"-install_name \\\\$rpath/\", \"-install_name @rpath/\")\n-            if self.settings.os == \"Android\" and tools.os_info.is_windows:\n-                # remove escape for quotation marks, to make ndk on windows happy\n-                tools.replace_in_file(\"configure\",\n-                    \"s/[\t ~#$^&*(){}\\\\\\\\|;'\\\\\\''\\\"<>?]/\\\\\\\\&/g\", \"s/[\t ~#$^&*(){}\\\\\\\\|;<>?]/\\\\\\\\&/g\")\n-        autotools = self._configure_autotools()\n-        autotools.make()\n-\n-    @functools.lru_cache(1)\n-    def _configure_autotools(self):\n-        autotools = AutoToolsBuildEnvironment(self, win_bash=tools.os_info.is_windows)\n-        autotools.libs = []\n-        yes_no = lambda v: \"yes\" if v else \"no\"\n-        args = [\n-            \"--enable-shared={}\".format(yes_no(self.options.shared)),\n-            \"--enable-static={}\".format(yes_no(not self.options.shared)),\n-        ]\n-        autotools.configure(args=args, configure_dir=self._source_subfolder)\n-        return autotools\n+        get(self, **self.conan_data[\"sources\"][self.version], strip_root=True)\n+\n+    def generate(self):\n+        if self._use_cmake:\n+            tc = CMakeToolchain(self)\n+            tc.variables[\"BUILD_PROGRAMS\"] = False\n+            tc.variables[\"FDK_AAC_INSTALL_CMAKE_CONFIG_MODULE\"] = False\n+            tc.variables[\"FDK_AAC_INSTALL_PKGCONFIG_MODULE\"] = False\n+            tc.generate()\n+        elif is_msvc(self):\n+            tc = NMakeToolchain(self)\n+            tc.generate()\n+        else:\n+            env = VirtualBuildEnv(self)\n+            env.generate()\n+            tc = AutotoolsToolchain(self)\n+            tc.generate()\n \n     def build(self):\n         if self._use_cmake:\n-            cmake = self._configure_cmake()\n+            cmake = CMake(self)\n+            cmake.configure()\n             cmake.build()\n-        elif self._is_msvc:\n-            self._build_vs()\n+        elif is_msvc(self):\n+            makefile_vc = os.path.join(self.source_folder, \"Makefile.vc\")\n+            replace_in_file(self, makefile_vc, \"CFLAGS   = /nologo /W3 /Ox /MT\", \"CFLAGS   = /nologo\")\n+            replace_in_file(self, makefile_vc, \"MKDIR_FLAGS = -p\", \"MKDIR_FLAGS =\")\n+            # Build either shared or static, and don't build utility (it always depends on static lib)\n+            replace_in_file(self, makefile_vc, \"copy $(PROGS) $(bindir)\", \"\")\n+            replace_in_file(self, makefile_vc, \"copy $(LIB_DEF) $(libdir)\", \"\")\n+            if self.options.shared:\n+                replace_in_file(\n+                    self, makefile_vc,\n+                    \"all: $(LIB_DEF) $(STATIC_LIB) $(SHARED_LIB) $(IMP_LIB) $(PROGS)\",\n+                    \"all: $(LIB_DEF) $(SHARED_LIB) $(IMP_LIB)\",\n+                )\n+                replace_in_file(self, makefile_vc, \"copy $(STATIC_LIB) $(libdir)\", \"\")\n+            else:\n+                replace_in_file(\n+                    self, makefile_vc,\n+                    \"all: $(LIB_DEF) $(STATIC_LIB) $(SHARED_LIB) $(IMP_LIB) $(PROGS)\",\n+                    \"all: $(STATIC_LIB)\",\n+                )\n+                replace_in_file(self, makefile_vc, \"copy $(IMP_LIB) $(libdir)\", \"\")\n+                replace_in_file(self, makefile_vc, \"copy $(SHARED_LIB) $(bindir)\", \"\")\n+            with chdir(self, self.source_folder):\n+                self.run(\"nmake -f Makefile.vc\")\n         else:\n-            self._build_autotools()\n+            autotools = Autotools(self)\n+            autotools.autoreconf()\n+            if self.settings.os == \"Android\" and self._settings_build.os == \"Windows\":\n+                # remove escape for quotation marks, to make ndk on windows happy\n+                replace_in_file(\n+                    self, os.path.join(self.source_folder, \"configure\"),\n+                    \"s/[\t ~#$^&*(){}\\\\\\\\|;'\\\\\\''\\\"<>?]/\\\\\\\\&/g\", \"s/[\t ~#$^&*(){}\\\\\\\\|;<>?]/\\\\\\\\&/g\",\n+                )\n+            autotools.configure()\n+            autotools.make()\n \n     def package(self):\n-        self.copy(pattern=\"NOTICE\", src=self._source_subfolder, dst=\"licenses\")\n+        copy(self, \"NOTICE\", src=self.source_folder, dst=os.path.join(self.package_folder, \"licenses\"))\n         if self._use_cmake:\n-            cmake = self._configure_cmake()\n+            cmake = CMake(self)\n             cmake.install()\n-        elif self._is_msvc:\n-            with self._msvc_build_environment():\n-                self.run(\"nmake -f Makefile.vc prefix=\\\"{}\\\" install\".format(self.package_folder))\n+        elif is_msvc(self):\n+            with chdir(self, self.source_folder):\n+                self.run(f\"nmake -f Makefile.vc prefix=\\\"{self.package_folder}\\\" install\")\n             if self.options.shared:\n-                tools.rename(os.path.join(self.package_folder, \"lib\", \"fdk-aac.dll.lib\"),\n+                rename(self, os.path.join(self.package_folder, \"lib\", \"fdk-aac.dll.lib\"),\n                              os.path.join(self.package_folder, \"lib\", \"fdk-aac.lib\"))\n         else:\n-            autotools = self._configure_autotools()\n+            autotools = Autotools(self)\n             autotools.install()\n-            tools.rmdir(os.path.join(self.package_folder, \"lib\", \"pkgconfig\"))\n-            tools.remove_files_by_mask(os.path.join(self.package_folder, \"lib\"), \"*.la\")\n+            rmdir(self, os.path.join(self.package_folder, \"lib\", \"pkgconfig\"))\n+            rm(self, \"*.la\", os.path.join(self.package_folder, \"lib\"))\n+            fix_apple_shared_install_name(self)\n \n     def package_info(self):\n         self.cpp_info.set_property(\"cmake_file_name\", \"fdk-aac\")\n@@ -167,7 +153,6 @@ def package_info(self):\n         self.cpp_info.filenames[\"cmake_find_package_multi\"] = \"fdk-aac\"\n         self.cpp_info.names[\"cmake_find_package\"] = \"FDK-AAC\"\n         self.cpp_info.names[\"cmake_find_package_multi\"] = \"FDK-AAC\"\n-        self.cpp_info.names[\"pkg_config\"] = \"fdk-aac\"\n         self.cpp_info.components[\"fdk-aac\"].names[\"cmake_find_package\"] = \"fdk-aac\"\n         self.cpp_info.components[\"fdk-aac\"].names[\"cmake_find_package_multi\"] = \"fdk-aac\"\n         self.cpp_info.components[\"fdk-aac\"].set_property(\"cmake_target_name\", \"FDK-AAC::fdk-aac\")"
		},
		{
		  "sha": "609976265e86c62b7c981b8ea63c987a57771111",
		  "filename": "recipes/libfdk_aac/all/test_package/CMakeLists.txt",
		  "status": "modified",
		  "additions": 4,
		  "deletions": 7,
		  "changes": 11,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/893ee53357b1914124b535113be20162e47f52cc/recipes%2Flibfdk_aac%2Fall%2Ftest_package%2FCMakeLists.txt",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/893ee53357b1914124b535113be20162e47f52cc/recipes%2Flibfdk_aac%2Fall%2Ftest_package%2FCMakeLists.txt",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Flibfdk_aac%2Fall%2Ftest_package%2FCMakeLists.txt?ref=893ee53357b1914124b535113be20162e47f52cc",
		  "patch": "@@ -1,11 +1,8 @@\n-cmake_minimum_required(VERSION 3.1)\n-project(test_package C)\n-\n-include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)\n-conan_basic_setup(TARGETS)\n+cmake_minimum_required(VERSION 3.8)\n+project(test_package LANGUAGES C)\n \n find_package(fdk-aac REQUIRED CONFIG)\n \n add_executable(${PROJECT_NAME} test_package.c)\n-target_link_libraries(${PROJECT_NAME} FDK-AAC::fdk-aac)\n-set_property(TARGET ${PROJECT_NAME} PROPERTY C_STANDARD 99)\n+target_link_libraries(${PROJECT_NAME} PRIVATE FDK-AAC::fdk-aac)\n+target_compile_features(${PROJECT_NAME} PRIVATE c_std_99)"
		},
		{
		  "sha": "0a6bc68712d90152ff3321a9468644f895d36c62",
		  "filename": "recipes/libfdk_aac/all/test_package/conanfile.py",
		  "status": "modified",
		  "additions": 14,
		  "deletions": 14,
		  "changes": 28,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/893ee53357b1914124b535113be20162e47f52cc/recipes%2Flibfdk_aac%2Fall%2Ftest_package%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/893ee53357b1914124b535113be20162e47f52cc/recipes%2Flibfdk_aac%2Fall%2Ftest_package%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Flibfdk_aac%2Fall%2Ftest_package%2Fconanfile.py?ref=893ee53357b1914124b535113be20162e47f52cc",
		  "patch": "@@ -1,26 +1,26 @@\n-from conans import ConanFile, CMake, tools\n+from conan import ConanFile\n+from conan.tools.build import can_run\n+from conan.tools.cmake import CMake, cmake_layout\n import os\n \n \n class TestPackageConan(ConanFile):\n-    settings = \"os\", \"compiler\", \"build_type\", \"arch\"\n-    generators = \"cmake\", \"cmake_find_package_multi\"\n+    settings = \"os\", \"arch\", \"compiler\", \"build_type\"\n+    generators = \"CMakeToolchain\", \"CMakeDeps\", \"VirtualRunEnv\"\n+    test_type = \"explicit\"\n \n-    def build_requirements(self):\n-        if self.settings.os == \"Macos\" and self.settings.arch == \"armv8\":\n-            # Workaround for CMake bug with error message:\n-            # Attempting to use @rpath without CMAKE_SHARED_LIBRARY_RUNTIME_C_FLAG being\n-            # set. This could be because you are using a Mac OS X version less than 10.5\n-            # or because CMake's platform configuration is corrupt.\n-            # FIXME: Remove once CMake on macOS/M1 CI runners is upgraded.\n-            self.build_requires(\"cmake/3.22.0\")\n+    def layout(self):\n+        cmake_layout(self)\n+\n+    def requirements(self):\n+        self.requires(self.tested_reference_str)\n \n     def build(self):\n         cmake = CMake(self)\n         cmake.configure()\n         cmake.build()\n \n     def test(self):\n-        if not tools.cross_building(self, skip_x64_x86=True):\n-            bin_path = os.path.join(\"bin\", \"test_package\")\n-            self.run(bin_path, run_environment=True)\n+        if can_run(self):\n+            bin_path = os.path.join(self.cpp.build.bindirs[0], \"test_package\")\n+            self.run(bin_path, env=\"conanrun\")"
		},
		{
		  "sha": "0d20897301b68bdd7b7c0a6fe54ad74b0e86c7f9",
		  "filename": "recipes/libfdk_aac/all/test_v1_package/CMakeLists.txt",
		  "status": "added",
		  "additions": 8,
		  "deletions": 0,
		  "changes": 8,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/893ee53357b1914124b535113be20162e47f52cc/recipes%2Flibfdk_aac%2Fall%2Ftest_v1_package%2FCMakeLists.txt",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/893ee53357b1914124b535113be20162e47f52cc/recipes%2Flibfdk_aac%2Fall%2Ftest_v1_package%2FCMakeLists.txt",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Flibfdk_aac%2Fall%2Ftest_v1_package%2FCMakeLists.txt?ref=893ee53357b1914124b535113be20162e47f52cc",
		  "patch": "@@ -0,0 +1,8 @@\n+cmake_minimum_required(VERSION 3.1)\n+project(test_package)\n+\n+include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)\n+conan_basic_setup(TARGETS)\n+\n+add_subdirectory(${CMAKE_CURRENT_SOURCE_DIR}/../test_package\n+                 ${CMAKE_CURRENT_BINARY_DIR}/test_package)"
		},
		{
		  "sha": "38f4483872d47f9327301676ba728e023b363c7f",
		  "filename": "recipes/libfdk_aac/all/test_v1_package/conanfile.py",
		  "status": "added",
		  "additions": 17,
		  "deletions": 0,
		  "changes": 17,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/893ee53357b1914124b535113be20162e47f52cc/recipes%2Flibfdk_aac%2Fall%2Ftest_v1_package%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/893ee53357b1914124b535113be20162e47f52cc/recipes%2Flibfdk_aac%2Fall%2Ftest_v1_package%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Flibfdk_aac%2Fall%2Ftest_v1_package%2Fconanfile.py?ref=893ee53357b1914124b535113be20162e47f52cc",
		  "patch": "@@ -0,0 +1,17 @@\n+from conans import ConanFile, CMake, tools\n+import os\n+\n+\n+class TestPackageConan(ConanFile):\n+    settings = \"os\", \"arch\", \"compiler\", \"build_type\"\n+    generators = \"cmake\", \"cmake_find_package_multi\"\n+\n+    def build(self):\n+        cmake = CMake(self)\n+        cmake.configure()\n+        cmake.build()\n+\n+    def test(self):\n+        if not tools.cross_building(self):\n+            bin_path = os.path.join(\"bin\", \"test_package\")\n+            self.run(bin_path, run_environment=True)"
		}
	  ]`)

	obtained, err := processChangedFiles(files)
	assert.Equal(t, err, nil)
	assert.Equal(t, &change{Recipe: "libfdk_aac", Change: EDIT, Weight: HEAVY}, obtained)
}

func TestProcessChangedFilesRegularNew(t *testing.T) {
	// https://github.com/conan-io/conan-center-index/pull/15260/files
	files := parsePullRequestFilesJSON(t, `[
		{
		  "sha": "632b160c486910b05c90767a8348b88797467233",
		  "filename": "recipes/kplot/all/CMakeLists.txt",
		  "status": "added",
		  "additions": 64,
		  "deletions": 0,
		  "changes": 64,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2FCMakeLists.txt",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2FCMakeLists.txt",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fkplot%2Fall%2FCMakeLists.txt?ref=ec5b98dc5f9030d3d33bd4b464857c560db20307",
		  "patch": "@@ -0,0 +1,64 @@\n+cmake_minimum_required(VERSION 3.8)\n+project(kplot LANGUAGES C)\n+\n+find_package(cairo REQUIRED CONFIG)\n+\n+set(kplot_src\n+    ${KPLOT_SRC_DIR}/colours.c\n+    ${KPLOT_SRC_DIR}/array.c\n+    ${KPLOT_SRC_DIR}/border.c\n+    ${KPLOT_SRC_DIR}/bucket.c\n+    ${KPLOT_SRC_DIR}/buffer.c\n+    ${KPLOT_SRC_DIR}/draw.c\n+    ${KPLOT_SRC_DIR}/grid.c\n+    ${KPLOT_SRC_DIR}/hist.c\n+    ${KPLOT_SRC_DIR}/label.c\n+    ${KPLOT_SRC_DIR}/kdata.c\n+    ${KPLOT_SRC_DIR}/kplot.c\n+    ${KPLOT_SRC_DIR}/margin.c\n+    ${KPLOT_SRC_DIR}/mean.c\n+    ${KPLOT_SRC_DIR}/plotctx.c\n+    ${KPLOT_SRC_DIR}/reallocarray.c\n+    ${KPLOT_SRC_DIR}/stddev.c\n+    ${KPLOT_SRC_DIR}/tic.c\n+    ${KPLOT_SRC_DIR}/vector.c\n+)\n+\n+set(kplot_inc\n+    ${KPLOT_SRC_DIR}/compat.h\n+    ${KPLOT_SRC_DIR}/extern.h\n+    ${KPLOT_SRC_DIR}/kplot.h\n+)\n+\n+include_directories(KPLOT_SRC_DIR)\n+\n+try_run(HAVE_reallocarray COMPIE_reallocarray ${CMAKE_BINARY_DIR} ${KPLOT_SRC_DIR}/test-reallocarray.c)\n+\n+file(READ \"${KPLOT_SRC_DIR}/compat.pre.h\" COMPAT_CONTENTS)\n+file(WRITE \"${KPLOT_SRC_DIR}/compat.h\" \"${COMPAT_CONTENTS}\")\n+if (${COMPIE_reallocarray} AND NOT ${HAVE_reallocarray})\n+    file(APPEND \"${KPLOT_SRC_DIR}/compat.h\" \"#define  HAVE_REALLOCARRAY\")\n+endif()\n+file(READ \"${KPLOT_SRC_DIR}/compat.post.h\" COMPAT_CONTENTS)\n+file(APPEND \"${KPLOT_SRC_DIR}/compat.h\" \"${COMPAT_CONTENTS}\")\n+\n+add_library(kplot ${kplot_src})\n+\n+target_compile_features(kplot PRIVATE c_std_99)\n+set_target_properties(kplot PROPERTIES\n+    PUBLIC_HEADER \"${kplot_inc}\"\n+    WINDOWS_EXPORT_ALL_SYMBOLS ON\n+    C_EXTENSIONS OFF\n+)\n+target_compile_features(kplot PRIVATE c_std_99)\n+target_link_libraries(kplot PRIVATE cairo::cairo)\n+\n+include(GNUInstallDirs)\n+\n+install(\n+    TARGETS kplot\n+    RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}\n+    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}\n+    ARCHIVE DESTINATION ${CMAKE_INSTALL_LIBDIR}\n+    PUBLIC_HEADER DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}\n+)"
		},
		{
		  "sha": "d0e63f3ee2c144b7b7ad2fe706eefea03f356264",
		  "filename": "recipes/kplot/all/conandata.yml",
		  "status": "added",
		  "additions": 4,
		  "deletions": 0,
		  "changes": 4,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Fconandata.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Fconandata.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fkplot%2Fall%2Fconandata.yml?ref=ec5b98dc5f9030d3d33bd4b464857c560db20307",
		  "patch": "@@ -0,0 +1,4 @@\n+sources:\n+  \"0.1.15\":\n+    url: \"https://github.com/kristapsdz/kplot/archive/refs/tags/VERSION_0_1_15.tar.gz\"\n+    sha256: \"602ebaac9b67dc7c7e84d8112df887c95ba0a1c4ed71fbab6671f8c5ecf4ba2a\""
		},
		{
		  "sha": "47e24a0f5cdf872f73cd082a6fb12610bf60c790",
		  "filename": "recipes/kplot/all/conanfile.py",
		  "status": "added",
		  "additions": 73,
		  "deletions": 0,
		  "changes": 73,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fkplot%2Fall%2Fconanfile.py?ref=ec5b98dc5f9030d3d33bd4b464857c560db20307",
		  "patch": "@@ -0,0 +1,73 @@\n+from conan import ConanFile\n+from conan.tools.files import get, copy\n+from conan.tools.cmake import CMake, CMakeDeps, CMakeToolchain, cmake_layout\n+from conan.tools.microsoft import is_msvc\n+from conan.errors import ConanInvalidConfiguration\n+import os\n+\n+required_conan_version = \">=1.53.0\"\n+\n+class KplotConan(ConanFile):\n+    name = \"kplot\"\n+    description = \"open source Cairo plotting library\"\n+    license = \"ISC\"\n+    url = \"https://github.com/conan-io/conan-center-index\"\n+    homepage = \"https://github.com/kristapsdz/kplot/\"\n+    topics = (\"plot\", \"cairo\", \"chart\") # no \"conan\"  and project name in topics\n+    settings = \"os\", \"arch\", \"compiler\", \"build_type\" # even for header only\n+    options = {\n+        \"shared\": [True, False],\n+        \"fPIC\": [True, False],\n+    }\n+    default_options = {\n+        \"shared\": False,\n+        \"fPIC\": True,\n+    }\n+\n+    def export_sources(self):\n+        copy(self, \"CMakeLists.txt\", src=self.recipe_folder, dst=self.export_sources_folder)\n+\n+    def config_options(self):\n+        if self.settings.os == \"Windows\":\n+            del self.options.fPIC\n+\n+    def configure(self):\n+        if self.options.shared:\n+            self.options.rm_safe(\"fPIC\")\n+        # for plain C projects only\n+        self.settings.rm_safe(\"compiler.libcxx\")\n+        self.settings.rm_safe(\"compiler.cppstd\")\n+\n+    def layout(self):\n+        cmake_layout(self, src_folder=\"src\")\n+\n+    def validate(self):\n+        if is_msvc(self):\n+            raise ConanInvalidConfiguration(f\"{self.ref} can not be built on Visual Studio and msvc.\")\n+\n+    def requirements(self):\n+        self.requires(\"cairo/1.17.4\")\n+\n+    def source(self):\n+        get(self, **self.conan_data[\"sources\"][self.version], strip_root=True)\n+\n+    def generate(self):\n+        tc = CMakeToolchain(self)\n+        tc.variables[\"KPLOT_SRC_DIR\"] = self.source_folder.replace(\"\\\\\", \"/\")\n+        tc.generate()\n+\n+        deps = CMakeDeps(self)\n+        deps.generate()\n+\n+    def build(self):\n+        cmake = CMake(self)\n+        cmake.configure(build_script_folder=os.path.join(self.source_folder, os.pardir))\n+        cmake.build()\n+\n+    def package(self):\n+        copy(self, pattern=\"LICENSE.md\", dst=os.path.join(self.package_folder, \"licenses\"), src=self.source_folder)\n+        cmake = CMake(self)\n+        cmake.install()\n+\n+    def package_info(self):\n+        self.cpp_info.libs = [\"kplot\"]"
		},
		{
		  "sha": "e61521b15e2af876876580af78cdefc3cf5b2f32",
		  "filename": "recipes/kplot/all/test_package/CMakeLists.txt",
		  "status": "added",
		  "additions": 8,
		  "deletions": 0,
		  "changes": 8,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Ftest_package%2FCMakeLists.txt",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Ftest_package%2FCMakeLists.txt",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fkplot%2Fall%2Ftest_package%2FCMakeLists.txt?ref=ec5b98dc5f9030d3d33bd4b464857c560db20307",
		  "patch": "@@ -0,0 +1,8 @@\n+cmake_minimum_required(VERSION 3.8)\n+project(test_package LANGUAGES C)\n+\n+find_package(kplot REQUIRED CONFIG)\n+\n+add_executable(${PROJECT_NAME} test_package.c)\n+target_link_libraries(${PROJECT_NAME} PRIVATE kplot::kplot)\n+target_compile_features(${PROJECT_NAME} PRIVATE c_std_99)"
		},
		{
		  "sha": "a9fb96656f2039c7269d31fb77077a2781551882",
		  "filename": "recipes/kplot/all/test_package/conanfile.py",
		  "status": "added",
		  "additions": 26,
		  "deletions": 0,
		  "changes": 26,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Ftest_package%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Ftest_package%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fkplot%2Fall%2Ftest_package%2Fconanfile.py?ref=ec5b98dc5f9030d3d33bd4b464857c560db20307",
		  "patch": "@@ -0,0 +1,26 @@\n+from conan import ConanFile\n+from conan.tools.build import can_run\n+from conan.tools.cmake import cmake_layout, CMake\n+import os\n+\n+\n+class TestPackageConan(ConanFile):\n+    settings = \"os\", \"arch\", \"compiler\", \"build_type\"\n+    generators = \"CMakeDeps\", \"CMakeToolchain\", \"VirtualRunEnv\"\n+    test_type = \"explicit\"\n+\n+    def requirements(self):\n+        self.requires(self.tested_reference_str)\n+\n+    def layout(self):\n+        cmake_layout(self)\n+\n+    def build(self):\n+        cmake = CMake(self)\n+        cmake.configure()\n+        cmake.build()\n+\n+    def test(self):\n+        if can_run(self):\n+            bin_path = os.path.join(self.cpp.build.bindirs[0], \"test_package\")\n+            self.run(bin_path, env=\"conanrun\")"
		},
		{
		  "sha": "2bc672875372a8e0bd23f06094638ef092af7840",
		  "filename": "recipes/kplot/all/test_package/test_package.c",
		  "status": "added",
		  "additions": 14,
		  "deletions": 0,
		  "changes": 14,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Ftest_package%2Ftest_package.c",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Ftest_package%2Ftest_package.c",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fkplot%2Fall%2Ftest_package%2Ftest_package.c?ref=ec5b98dc5f9030d3d33bd4b464857c560db20307",
		  "patch": "@@ -0,0 +1,14 @@\n+#include <stdlib.h>\n+#include <unistd.h>\n+\n+#include \"cairo.h\"\n+#include \"kplot.h\"\n+\n+int main() {\n+    struct kpair points1[50];\n+    struct kdata* d1 = kdata_array_alloc(points1, 50);\n+\n+    kdata_destroy(d1);\n+\n+    return 0;\n+}"
		},
		{
		  "sha": "925ecbe19e448d148f983a2ba5df4e55ebe76904",
		  "filename": "recipes/kplot/all/test_v1_package/CMakeLists.txt",
		  "status": "added",
		  "additions": 8,
		  "deletions": 0,
		  "changes": 8,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Ftest_v1_package%2FCMakeLists.txt",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Ftest_v1_package%2FCMakeLists.txt",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fkplot%2Fall%2Ftest_v1_package%2FCMakeLists.txt?ref=ec5b98dc5f9030d3d33bd4b464857c560db20307",
		  "patch": "@@ -0,0 +1,8 @@\n+cmake_minimum_required(VERSION 3.1)\n+project(test_package)\n+\n+include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)\n+conan_basic_setup(TARGETS)\n+\n+add_subdirectory(${CMAKE_CURRENT_SOURCE_DIR}/../test_package/\n+                 ${CMAKE_CURRENT_BINARY_DIR}/test_package/)"
		},
		{
		  "sha": "5a05af3c2dfd2f512de62e66004863e5e13d5d90",
		  "filename": "recipes/kplot/all/test_v1_package/conanfile.py",
		  "status": "added",
		  "additions": 18,
		  "deletions": 0,
		  "changes": 18,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Ftest_v1_package%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fall%2Ftest_v1_package%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fkplot%2Fall%2Ftest_v1_package%2Fconanfile.py?ref=ec5b98dc5f9030d3d33bd4b464857c560db20307",
		  "patch": "@@ -0,0 +1,18 @@\n+from conans import ConanFile, CMake\n+from conan.tools.build import cross_building\n+import os\n+\n+\n+class TestPackageV1Conan(ConanFile):\n+    settings = \"os\", \"arch\", \"compiler\", \"build_type\"\n+    generators = \"cmake\", \"cmake_find_package_multi\"\n+\n+    def build(self):\n+        cmake = CMake(self)\n+        cmake.configure()\n+        cmake.build()\n+\n+    def test(self):\n+        if not cross_building(self):\n+            bin_path = os.path.join(\"bin\", \"test_package\")\n+            self.run(bin_path, run_environment=True)"
		},
		{
		  "sha": "d90a2c03c7832b8771fe546670003dcf4416e434",
		  "filename": "recipes/kplot/config.yml",
		  "status": "added",
		  "additions": 3,
		  "deletions": 0,
		  "changes": 3,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fconfig.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/ec5b98dc5f9030d3d33bd4b464857c560db20307/recipes%2Fkplot%2Fconfig.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fkplot%2Fconfig.yml?ref=ec5b98dc5f9030d3d33bd4b464857c560db20307",
		  "patch": "@@ -0,0 +1,3 @@\n+versions:\n+  \"0.1.15\":\n+    folder: all"
		}
	  ]`)

	obtained, err := processChangedFiles(files)
	assert.Equal(t, err, nil)
	assert.Equal(t, &change{Recipe: "kplot", Change: NEW, Weight: REGULAR}, obtained)
}

func TestProcessChangedFileRegularEditTwo(t *testing.T) {
	// https://github.com/conan-io/conan-center-index/pull/15293/files
	files := parsePullRequestFilesJSON(t, `[
  {
    "sha": "85ea4e8fd6a622f87d8ac0419acf13e5f12fec86",
    "filename": "recipes/spix/all/conandata.yml",
    "status": "modified",
    "additions": 7,
    "deletions": 0,
    "changes": 7,
    "blob_url": "https://github.com/conan-io/conan-center-index/blob/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Fconandata.yml",
    "raw_url": "https://github.com/conan-io/conan-center-index/raw/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Fconandata.yml",
    "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fspix%2Fall%2Fconandata.yml?ref=dcd594527464ad665fc52825a8daa9ff3607b270",
    "patch": "@@ -2,8 +2,15 @@ sources:\n   \"0.4\":\n     url: \"https://github.com/faaxm/spix/archive/refs/tags/v0.4.tar.gz\"\n     sha256: \"e787c08840c37e5b153c0139f3bb613a2729ae8f6ccd0fb450fef92971cd8b53\"\n+  \"0.5\":\n+    url: \"https://github.com/faaxm/spix/archive/refs/tags/v0.5.tar.gz\"\n+    sha256: \"d3fd9bb069aef6ff6c93c69524ed3603afd24e6b52e4bb8d093c80cec255d4dc\"\n patches:\n   \"0.4\":\n     - patch_file: \"patches/0001-use-conan-libs-0.4.patch\"\n       patch_description: \"Link to conan libs\"\n       patch_type: \"conan\"\n+  \"0.5\":\n+    - patch_file: \"patches/0001-use-conan-libs-0.5.patch\"\n+      patch_description: \"Link to conan libs\"\n+      patch_type: \"conan\""
  },
  {
    "sha": "62b8b81ba14100f8517aa83ec8563dcc9525825b",
    "filename": "recipes/spix/all/conanfile.py",
    "status": "modified",
    "additions": 17,
    "deletions": 10,
    "changes": 27,
    "blob_url": "https://github.com/conan-io/conan-center-index/blob/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Fconanfile.py",
    "raw_url": "https://github.com/conan-io/conan-center-index/raw/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Fconanfile.py",
    "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fspix%2Fall%2Fconanfile.py?ref=dcd594527464ad665fc52825a8daa9ff3607b270",
    "patch": "@@ -29,16 +29,24 @@ class SpixConan(ConanFile):\n \n     @property\n     def _minimum_cpp_standard(self):\n-        return 14\n+        return 14 if self.version == \"0.4\" else 17\n \n     @property\n     def _compilers_minimum_version(self):\n-        return {\n-            \"Visual Studio\": \"14\",\n-            \"gcc\": \"5\",\n-            \"clang\": \"3.4\",\n-            \"apple-clang\": \"10\"\n-        }\n+        if self.version == \"0.4\":\n+            return {\n+                \"Visual Studio\": \"14\",\n+                \"gcc\": \"5\",\n+                \"clang\": \"3.4\",\n+                \"apple-clang\": \"10\"\n+            }\n+        else:\n+            return {\n+                \"Visual Studio\": \"15.7\",\n+                \"gcc\": \"7\",\n+                \"clang\": \"5\",\n+                \"apple-clang\": \"10\",\n+            }\n \n     def export_sources(self):\n         export_conandata_patches(self)\n@@ -59,8 +67,7 @@ def layout(self):\n \n     def requirements(self):\n         self.requires(\"anyrpc/1.0.2\")\n-        self.requires(\"qt/6.3.1\")\n-        self.requires(\"expat/2.4.9\")\n+        self.requires(\"qt/6.4.2\")\n         \n     def validate(self):\n         if self.info.settings.compiler.cppstd:\n@@ -91,7 +98,7 @@ def generate(self):\n \n     def _patch_sources(self):\n         apply_conandata_patches(self)\n-        if Version(self.deps_cpp_info[\"qt\"].version).major == 6:\n+        if self.version == \"0.4\" and Version(self.deps_cpp_info[\"qt\"].version).major == 6:\n             replace_in_file(self, os.path.join(self.source_folder, \"CMakeLists.txt\"), \"set(CMAKE_CXX_STANDARD 14)\", \"set(CMAKE_CXX_STANDARD 17)\")\n \n     def build(self):"
  },
  {
    "sha": "079607701b2fb98b0de72c79843dbd13fbfa0f9c",
    "filename": "recipes/spix/all/patches/0001-use-conan-libs-0.5.patch",
    "status": "added",
    "additions": 32,
    "deletions": 0,
    "changes": 32,
    "blob_url": "https://github.com/conan-io/conan-center-index/blob/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Fpatches%2F0001-use-conan-libs-0.5.patch",
    "raw_url": "https://github.com/conan-io/conan-center-index/raw/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Fpatches%2F0001-use-conan-libs-0.5.patch",
    "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fspix%2Fall%2Fpatches%2F0001-use-conan-libs-0.5.patch?ref=dcd594527464ad665fc52825a8daa9ff3607b270",
    "patch": "@@ -0,0 +1,32 @@\n+--- a/CMakeLists.txt\n++++ b/CMakeLists.txt\n+@@ -5,7 +5,6 @@ option(SPIX_BUILD_EXAMPLES \"Build Spix examples.\" ON)\n+ option(SPIX_BUILD_TESTS \"Build Spix unit tests.\" OFF)\n+ set(SPIX_QT_MAJOR \"6\" CACHE STRING \"Major Qt version to build Spix against\")\n+ \n+-set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} \"${CMAKE_CURRENT_LIST_DIR}/cmake/modules\")\n+ set(CMAKE_CXX_STANDARD 17)\n+ \n+ # Hide symbols unless explicitly flagged with SPIX_EXPORT\n+diff --git a/lib/CMakeLists.txt b/lib/CMakeLists.txt\n+index 723de5e..f234bec 100644\n+--- a/lib/CMakeLists.txt\n++++ b/lib/CMakeLists.txt\n+@@ -8,7 +8,7 @@ include(CMakePackageConfigHelpers)\n+ # Dependencies\n+ #\n+ find_package(Threads REQUIRED)\n+-find_package(AnyRPC REQUIRED)\n++find_package(anyrpc REQUIRED)\n+ find_package(Qt${SPIX_QT_MAJOR}\n+     COMPONENTS\n+         Core\n+@@ -132,7 +132,7 @@ target_link_libraries(Spix\n+         Qt${SPIX_QT_MAJOR}::Gui\n+         Qt${SPIX_QT_MAJOR}::Quick\n+     PRIVATE\n+-        AnyRPC::anyrpc\n++        anyrpc::anyrpc\n+ )\n+ \n+ #"
  },
  {
    "sha": "8b674ce10f424662cc3f7aed09bf75be52a74a21",
    "filename": "recipes/spix/all/test_package/CMakeLists.txt",
    "status": "modified",
    "additions": 1,
    "deletions": 1,
    "changes": 2,
    "blob_url": "https://github.com/conan-io/conan-center-index/blob/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Ftest_package%2FCMakeLists.txt",
    "raw_url": "https://github.com/conan-io/conan-center-index/raw/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Ftest_package%2FCMakeLists.txt",
    "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fspix%2Fall%2Ftest_package%2FCMakeLists.txt?ref=dcd594527464ad665fc52825a8daa9ff3607b270",
    "patch": "@@ -6,4 +6,4 @@ find_package(Spix REQUIRED CONFIG)\n \n add_executable(${PROJECT_NAME} test_spix.cpp)\n target_link_libraries(${PROJECT_NAME} PRIVATE Spix::Spix)\n-target_compile_features(${PROJECT_NAME} PRIVATE cxx_std_14)\n+target_compile_features(${PROJECT_NAME} PRIVATE cxx_std_17)"
  },
  {
    "sha": "b47e9bb0e34131b0caf5d1846fda0ca05cf6a80e",
    "filename": "recipes/spix/all/test_package/conanfile.py",
    "status": "modified",
    "additions": 0,
    "deletions": 5,
    "changes": 5,
    "blob_url": "https://github.com/conan-io/conan-center-index/blob/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Ftest_package%2Fconanfile.py",
    "raw_url": "https://github.com/conan-io/conan-center-index/raw/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Ftest_package%2Fconanfile.py",
    "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fspix%2Fall%2Ftest_package%2Fconanfile.py?ref=dcd594527464ad665fc52825a8daa9ff3607b270",
    "patch": "@@ -17,12 +17,7 @@ def requirements(self):\n     def layout(self):\n         cmake_layout(self)\n \n-    def _patch_sources(self):\n-        if Version(self.deps_cpp_info[\"qt\"].version).major == 6:\n-            replace_in_file(self, os.path.join(self.source_folder, \"CMakeLists.txt\"), \"cxx_std_14\", \"cxx_std_17\")\n-\n     def build(self):\n-        self._patch_sources()\n         cmake = CMake(self)\n         cmake.configure()\n         cmake.build()"
  },
  {
    "sha": "131e8cc2de764dfc59b4e0041f0110f2b7db5ec7",
    "filename": "recipes/spix/all/test_v1_package/conanfile.py",
    "status": "modified",
    "additions": 0,
    "deletions": 5,
    "changes": 5,
    "blob_url": "https://github.com/conan-io/conan-center-index/blob/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Ftest_v1_package%2Fconanfile.py",
    "raw_url": "https://github.com/conan-io/conan-center-index/raw/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fall%2Ftest_v1_package%2Fconanfile.py",
    "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fspix%2Fall%2Ftest_v1_package%2Fconanfile.py?ref=dcd594527464ad665fc52825a8daa9ff3607b270",
    "patch": "@@ -9,12 +9,7 @@ class TestSpixV1Conan(ConanFile):\n     settings = \"os\", \"arch\", \"compiler\", \"build_type\"\n     generators = \"cmake\", \"cmake_find_package_multi\"\n \n-    def _patch_sources(self):\n-        if Version(self.deps_cpp_info[\"qt\"].version).major == 6:\n-            replace_in_file(self, os.path.join(self.source_folder, \"CMakeLists.txt\"), \"cxx_std_14\", \"cxx_std_17\")\n-\n     def build(self):\n-        self._patch_sources()\n         cmake = CMake(self)\n         cmake.configure()\n         cmake.build()"
  },
  {
    "sha": "57a887729c2abb62f3b5772102328b71fd6c5078",
    "filename": "recipes/spix/config.yml",
    "status": "modified",
    "additions": 2,
    "deletions": 0,
    "changes": 2,
    "blob_url": "https://github.com/conan-io/conan-center-index/blob/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fconfig.yml",
    "raw_url": "https://github.com/conan-io/conan-center-index/raw/dcd594527464ad665fc52825a8daa9ff3607b270/recipes%2Fspix%2Fconfig.yml",
    "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fspix%2Fconfig.yml?ref=dcd594527464ad665fc52825a8daa9ff3607b270",
    "patch": "@@ -1,3 +1,5 @@\n versions:\n   \"0.4\":\n     folder: all\n+  \"0.5\":\n+    folder: all"
  }
]`)

	obtained, err := processChangedFiles(files)
	assert.Equal(t, err, nil)
	assert.Equal(t, &change{Recipe: "spix", Change: EDIT, Weight: REGULAR}, obtained)
}

func TestProcessChangedFilesSmallEditOne(t *testing.T) {
	// https://github.com/conan-io/conan-center-index/pull/15416/files
	files := parsePullRequestFilesJSON(t, `[
		{
		  "sha": "1da6086070fb7096513923956f5ac8331b2f4fa1",
		  "filename": "recipes/trantor/all/conandata.yml",
		  "status": "modified",
		  "additions": 3,
		  "deletions": 0,
		  "changes": 3,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/0095be1f737baaa849ddf1ee2f92304eba965b4a/recipes%2Ftrantor%2Fall%2Fconandata.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/0095be1f737baaa849ddf1ee2f92304eba965b4a/recipes%2Ftrantor%2Fall%2Fconandata.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Ftrantor%2Fall%2Fconandata.yml?ref=0095be1f737baaa849ddf1ee2f92304eba965b4a",
		  "patch": "@@ -1,4 +1,7 @@\n sources:\n+  \"1.5.10\":\n+    url: \"https://github.com/an-tao/trantor/archive/v1.5.10.tar.gz\"\n+    sha256: \"2d47775b3091a1a103bea46f5da017dc03c39883f8d717cf6ba24bdcdf01a15d\"\n   \"1.5.8\":\n     url: \"https://github.com/an-tao/trantor/archive/v1.5.8.tar.gz\"\n     sha256: \"705ec0176681be5c99fcc7af37416ece9d65ff4d907bca764cb11471b104fbf8\""
		},
		{
		  "sha": "ceac8f815ba356008c54c4efce715eeaa7fd0a3f",
		  "filename": "recipes/trantor/all/conanfile.py",
		  "status": "modified",
		  "additions": 4,
		  "deletions": 6,
		  "changes": 10,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/0095be1f737baaa849ddf1ee2f92304eba965b4a/recipes%2Ftrantor%2Fall%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/0095be1f737baaa849ddf1ee2f92304eba965b4a/recipes%2Ftrantor%2Fall%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Ftrantor%2Fall%2Fconanfile.py?ref=0095be1f737baaa849ddf1ee2f92304eba965b4a",
		  "patch": "@@ -8,7 +8,7 @@\n \n import os\n \n-required_conan_version = \">=1.52.0\"\n+required_conan_version = \">=1.53.0\"\n \n class TrantorConan(ConanFile):\n     name = \"trantor\"\n@@ -38,6 +38,7 @@ def _compilers_minimum_version(self):\n         return {\n             \"gcc\": \"5\",\n             \"Visual Studio\": \"15.0\",\n+            \"msvc\": \"191\",\n             \"clang\": \"5\",\n             \"apple-clang\": \"10\",\n         }\n@@ -48,18 +49,15 @@ def config_options(self):\n \n     def configure(self):\n         if self.options.shared:\n-            try:\n-                del self.options.fPIC\n-            except Exception:\n-                pass\n+            self.options.rm_safe(\"fPIC\")\n \n     def layout(self):\n         cmake_layout(self, src_folder=\"src\")\n \n     def requirements(self):\n         self.requires(\"openssl/1.1.1s\")\n         if self.options.with_c_ares:\n-            self.requires(\"c-ares/1.18.1\")\n+            self.requires(\"c-ares/1.19.0\")\n \n     def validate(self):\n         if self.info.settings.compiler.get_safe(\"cppstd\"):"
		},
		{
		  "sha": "96e466512e5b23e441aa51b8bc65924a56c1a49f",
		  "filename": "recipes/trantor/all/test_package/CMakeLists.txt",
		  "status": "modified",
		  "additions": 2,
		  "deletions": 2,
		  "changes": 4,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/0095be1f737baaa849ddf1ee2f92304eba965b4a/recipes%2Ftrantor%2Fall%2Ftest_package%2FCMakeLists.txt",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/0095be1f737baaa849ddf1ee2f92304eba965b4a/recipes%2Ftrantor%2Fall%2Ftest_package%2FCMakeLists.txt",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Ftrantor%2Fall%2Ftest_package%2FCMakeLists.txt?ref=0095be1f737baaa849ddf1ee2f92304eba965b4a",
		  "patch": "@@ -1,8 +1,8 @@\n cmake_minimum_required(VERSION 3.8)\n-project(test_package CXX)\n+project(test_package LANGUAGES CXX)\n \n find_package(Trantor CONFIG REQUIRED)\n \n add_executable(${PROJECT_NAME} test_package.cpp)\n-target_link_libraries(${PROJECT_NAME} Trantor::Trantor)\n+target_link_libraries(${PROJECT_NAME} PRIVATE Trantor::Trantor)\n target_compile_features(${PROJECT_NAME} PRIVATE cxx_std_14)"
		},
		{
		  "sha": "bc541ea90b5128a1b011f64ea03ec20f030c49bd",
		  "filename": "recipes/trantor/all/test_v1_package/CMakeLists.txt",
		  "status": "modified",
		  "additions": 3,
		  "deletions": 6,
		  "changes": 9,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/0095be1f737baaa849ddf1ee2f92304eba965b4a/recipes%2Ftrantor%2Fall%2Ftest_v1_package%2FCMakeLists.txt",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/0095be1f737baaa849ddf1ee2f92304eba965b4a/recipes%2Ftrantor%2Fall%2Ftest_v1_package%2FCMakeLists.txt",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Ftrantor%2Fall%2Ftest_v1_package%2FCMakeLists.txt?ref=0095be1f737baaa849ddf1ee2f92304eba965b4a",
		  "patch": "@@ -1,12 +1,9 @@\n cmake_minimum_required(VERSION 3.8)\n \n-project(test_package CXX)\n+project(test_package)\n \n include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)\n conan_basic_setup(TARGETS)\n \n-find_package(Trantor REQUIRED CONFIG)\n-\n-add_executable(${PROJECT_NAME} ../test_package/test_package.cpp)\n-target_link_libraries(${PROJECT_NAME} PRIVATE Trantor::Trantor)\n-target_compile_features(${PROJECT_NAME} PRIVATE cxx_std_14)\n+add_subdirectory(${CMAKE_CURRENT_SOURCE_DIR}/../test_package/\n+                 ${CMAKE_CURRENT_BINARY_DIR}/test_package/)"
		},
		{
		  "sha": "0056a4f6c8a1d2df8ea999c6eda50efbaba0a1fd",
		  "filename": "recipes/trantor/config.yml",
		  "status": "modified",
		  "additions": 2,
		  "deletions": 0,
		  "changes": 2,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/0095be1f737baaa849ddf1ee2f92304eba965b4a/recipes%2Ftrantor%2Fconfig.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/0095be1f737baaa849ddf1ee2f92304eba965b4a/recipes%2Ftrantor%2Fconfig.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Ftrantor%2Fconfig.yml?ref=0095be1f737baaa849ddf1ee2f92304eba965b4a",
		  "patch": "@@ -1,4 +1,6 @@\n versions:\n+  \"1.5.10\":\n+    folder: \"all\"\n   \"1.5.8\":\n     folder: \"all\"\n   \"1.5.7\":"
		}
	  ]`)

	obtained, err := processChangedFiles(files)
	assert.Equal(t, err, nil)
	assert.Equal(t, &change{Recipe: "trantor", Change: EDIT, Weight: SMALL}, obtained)
}

func TestProcessChangedFilesRegularEditTwo(t *testing.T) {
	// https://github.com/conan-io/conan-center-index/pull/15594/files
	files := parsePullRequestFilesJSON(t, `[
		{
		  "sha": "ac9853ed7cd15facd13379a89509be9f834e2d18",
		  "filename": "recipes/libwebp/all/conandata.yml",
		  "status": "modified",
		  "additions": 16,
		  "deletions": 16,
		  "changes": 32,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/e811994dc59acac393c90b255009548bafb99dcc/recipes%2Flibwebp%2Fall%2Fconandata.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/e811994dc59acac393c90b255009548bafb99dcc/recipes%2Flibwebp%2Fall%2Fconandata.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Flibwebp%2Fall%2Fconandata.yml?ref=e811994dc59acac393c90b255009548bafb99dcc",
		  "patch": "@@ -1,28 +1,28 @@\n sources:\n   \"1.3.0\":\n-    url: \"https://github.com/webmproject/libwebp/archive/v1.3.0.tar.gz\"\n-    sha256: \"dc9860d3fe06013266c237959e1416b71c63b36f343aae1d65ea9c94832630e1\"\n+    url: \"https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.3.0.tar.gz\"\n+    sha256: \"64ac4614db292ae8c5aa26de0295bf1623dbb3985054cb656c55e67431def17c\"\n   \"1.2.4\":\n-    url: \"https://github.com/webmproject/libwebp/archive/v1.2.4.tar.gz\"\n-    sha256: \"dfe7bff3390cd4958da11e760b65318f0a48c32913e4d5bc5e8d55abaaa2d32e\"\n+    url: \"https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.2.4.tar.gz\"\n+    sha256: \"7bf5a8a28cc69bcfa8cb214f2c3095703c6b73ac5fba4d5480c205331d9494df\"\n   \"1.2.3\":\n-    url: \"https://github.com/webmproject/libwebp/archive/v1.2.3.tar.gz\"\n-    sha256: \"021169407825d7ad918ff4554c6af885e7cf116d9b641cfd7f04c1173ffb9eb0\"\n+    url: \"https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.2.3.tar.gz\"\n+    sha256: \"f5d7ab2390b06b8a934a4fc35784291b3885b557780d099bd32f09241f9d83f9\"\n   \"1.2.2\":\n-    url: \"https://github.com/webmproject/libwebp/archive/v1.2.2.tar.gz\"\n-    sha256: \"51e9297aadb7d9eb99129fe0050f53a11fcce38a0848fb2b0389e385ad93695e\"\n+    url: \"https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.2.2.tar.gz\"\n+    sha256: \"7656532f837af5f4cec3ff6bafe552c044dc39bf453587bd5b77450802f4aee6\"\n   \"1.2.1\":\n-    url: \"https://github.com/webmproject/libwebp/archive/v1.2.1.tar.gz\"\n-    sha256: \"01bcde6a40a602294994050b81df379d71c40b7e39c819c024d079b3c56307f4\"\n+    url: \"https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.2.1.tar.gz\"\n+    sha256: \"808b98d2f5b84e9b27fdef6c5372dac769c3bda4502febbfa5031bd3c4d7d018\"\n   \"1.2.0\":\n-    url: \"https://github.com/webmproject/libwebp/archive/v1.2.0.tar.gz\"\n-    sha256: \"d60608c45682fa1e5d41c3c26c199be5d0184084cd8a971a6fc54035f76487d3\"\n+    url: \"https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.2.0.tar.gz\"\n+    sha256: \"2fc8bbde9f97f2ab403c0224fb9ca62b2e6852cbc519e91ceaa7c153ffd88a0c\"\n   \"1.1.0\":\n-    url: \"https://github.com/webmproject/libwebp/archive/v1.1.0.tar.gz\"\n-    sha256: \"424faab60a14cb92c2a062733b6977b4cc1e875a6398887c5911b3a1a6c56c51\"\n+    url: \"https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.1.0.tar.gz\"\n+    sha256: \"98a052268cc4d5ece27f76572a7f50293f439c17a98e67c4ea0c7ed6f50ef043\"\n   \"1.0.3\":\n-    url: \"https://github.com/webmproject/libwebp/archive/v1.0.3.tar.gz\"\n-    sha256: \"082d114bcb18a0e2aafc3148d43367c39304f86bf18ba0b2e766447e111a4a91\"\n+    url: \"https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.0.3.tar.gz\"\n+    sha256: \"e20a07865c8697bba00aebccc6f54912d6bc333bb4d604e6b07491c1a226b34f\"\n patches:\n   \"1.3.0\":\n     - patch_file: \"patches/1.3.0-0001-fix-cmake.patch\""
		},
		{
		  "sha": "8e1d3bc0866e49a89c16adb8c512007c88ff0185",
		  "filename": "recipes/libwebp/all/conanfile.py",
		  "status": "modified",
		  "additions": 4,
		  "deletions": 5,
		  "changes": 9,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/e811994dc59acac393c90b255009548bafb99dcc/recipes%2Flibwebp%2Fall%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/e811994dc59acac393c90b255009548bafb99dcc/recipes%2Flibwebp%2Fall%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Flibwebp%2Fall%2Fconanfile.py?ref=e811994dc59acac393c90b255009548bafb99dcc",
		  "patch": "@@ -12,7 +12,7 @@ class LibwebpConan(ConanFile):\n     name = \"libwebp\"\n     description = \"Library to encode and decode images in WebP format\"\n     url = \"https://github.com/conan-io/conan-center-index\"\n-    homepage = \"https://github.com/webmproject/libwebp\"\n+    homepage = \"https://chromium.googlesource.com/webm/libwebp\"\n     topics = (\"image\", \"libwebp\", \"webp\", \"decoding\", \"encoding\")\n     license = \"BSD-3-Clause\"\n \n@@ -49,8 +49,7 @@ def layout(self):\n         cmake_layout(self, src_folder=\"src\")\n \n     def source(self):\n-        get(self, **self.conan_data[\"sources\"][self.version],\n-            destination=self.source_folder, strip_root=True)\n+        get(self, **self.conan_data[\"sources\"][self.version], strip_root=True)\n \n     def generate(self):\n         tc = CMakeToolchain(self)\n@@ -78,8 +77,8 @@ def generate(self):\n             tc.variables[\"WEBP_BUILD_LIBWEBPMUX\"] = True\n         tc.variables[\"WEBP_BUILD_WEBPMUX\"] = False\n         if self.options.shared and is_msvc(self):\n-          # Building a dll (see fix-dll-export patch)\n-          tc.preprocessor_definitions[\"WEBP_DLL\"] = 1\n+            # Building a dll (see fix-dll-export patch)\n+            tc.preprocessor_definitions[\"WEBP_DLL\"] = 1\n         tc.generate()\n \n     def build(self):"
		}
	  ]`)

	obtained, err := processChangedFiles(files)
	assert.Equal(t, err, nil)
	assert.Equal(t, &change{Recipe: "libwebp", Change: EDIT, Weight: REGULAR}, obtained)
}

func TestProcessChangedFilesRegularEditThree(t *testing.T) {
	// https://github.com/conan-io/conan-center-index/pull/15470/files
	files := parsePullRequestFilesJSON(t, `[
		{
		  "sha": "c700cd876b542bca99f5aec953554e406597e148",
		  "filename": "recipes/cpp-peglib/1.x.x/conanfile.py",
		  "status": "modified",
		  "additions": 10,
		  "deletions": 13,
		  "changes": 23,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/08c0d0373f5008624a472098615d5de48d056c0a/recipes%2Fcpp-peglib%2F1.x.x%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/08c0d0373f5008624a472098615d5de48d056c0a/recipes%2Fcpp-peglib%2F1.x.x%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fcpp-peglib%2F1.x.x%2Fconanfile.py?ref=08c0d0373f5008624a472098615d5de48d056c0a",
		  "patch": "@@ -1,14 +1,13 @@\n from conan import ConanFile\n from conan.errors import ConanInvalidConfiguration\n-from conan.tools.build import check_min_cppstd\n+from conan.tools.build import check_min_cppstd, stdcpp_library\n from conan.tools.files import copy, get\n from conan.tools.layout import basic_layout\n from conan.tools.microsoft import is_msvc\n from conan.tools.scm import Version\n-from conans import tools as tools_legacy\n import os\n \n-required_conan_version = \">=1.50.0\"\n+required_conan_version = \">=1.54.0\"\n \n \n class CpppeglibConan(ConanFile):\n@@ -18,6 +17,7 @@ class CpppeglibConan(ConanFile):\n     url = \"https://github.com/conan-io/conan-center-index\"\n     homepage = \"https://github.com/yhirose/cpp-peglib\"\n     topics = (\"peg\", \"parser\", \"header-only\")\n+    package_type = \"header-library\"\n     settings = \"os\", \"arch\", \"compiler\", \"build_type\"\n     no_copy_source = True\n \n@@ -31,9 +31,12 @@ def _compilers_minimum_version(self):\n             \"Visual Studio\": \"15.7\",\n             \"gcc\": \"7\",\n             \"clang\": \"6\",\n-            \"apple-clang\": \"10\"\n+            \"apple-clang\": \"10\",\n         }\n \n+    def layout(self):\n+        basic_layout(self, src_folder=\"src\")\n+\n     def package_id(self):\n         self.info.clear()\n \n@@ -50,19 +53,15 @@ def loose_lt_semver(v1, v2):\n         minimum_version = self._compilers_minimum_version.get(str(self.settings.compiler), False)\n         if minimum_version and loose_lt_semver(str(self.settings.compiler.version), minimum_version):\n             raise ConanInvalidConfiguration(\n-                f\"{self.name} {self.version} requires C++{self._min_cppstd}, which your compiler does not support.\",\n+                f\"{self.ref} requires C++{self._min_cppstd}, which your compiler does not support.\",\n             )\n \n         if self.settings.compiler == \"clang\" and Version(self.settings.compiler.version) == \"7\" and \\\n-           tools_legacy.stdcpp_library(self) == \"stdc++\":\n+           stdcpp_library(self) == \"stdc++\":\n             raise ConanInvalidConfiguration(f\"{self.name} {self.version} does not support clang 7 with libstdc++.\")\n \n-    def layout(self):\n-        basic_layout(self, src_folder=\"src\")\n-\n     def source(self):\n-        get(self, **self.conan_data[\"sources\"][self.version],\n-            destination=self.source_folder, strip_root=True)\n+        get(self, **self.conan_data[\"sources\"][self.version], strip_root=True)\n \n     def build(self):\n         pass\n@@ -73,9 +72,7 @@ def package(self):\n \n     def package_info(self):\n         self.cpp_info.bindirs = []\n-        self.cpp_info.frameworkdirs = []\n         self.cpp_info.libdirs = []\n-        self.cpp_info.resdirs = []\n         if self.settings.os in [\"Linux\", \"FreeBSD\"]:\n             self.cpp_info.system_libs = [\"pthread\"]\n             self.cpp_info.cxxflags.append(\"-pthread\")"
		},
		{
		  "sha": "0a6bc68712d90152ff3321a9468644f895d36c62",
		  "filename": "recipes/cpp-peglib/1.x.x/test_package/conanfile.py",
		  "status": "modified",
		  "additions": 4,
		  "deletions": 3,
		  "changes": 7,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/08c0d0373f5008624a472098615d5de48d056c0a/recipes%2Fcpp-peglib%2F1.x.x%2Ftest_package%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/08c0d0373f5008624a472098615d5de48d056c0a/recipes%2Fcpp-peglib%2F1.x.x%2Ftest_package%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fcpp-peglib%2F1.x.x%2Ftest_package%2Fconanfile.py?ref=08c0d0373f5008624a472098615d5de48d056c0a",
		  "patch": "@@ -7,13 +7,14 @@\n class TestPackageConan(ConanFile):\n     settings = \"os\", \"arch\", \"compiler\", \"build_type\"\n     generators = \"CMakeToolchain\", \"CMakeDeps\", \"VirtualRunEnv\"\n-\n-    def requirements(self):\n-        self.requires(self.tested_reference_str)\n+    test_type = \"explicit\"\n \n     def layout(self):\n         cmake_layout(self)\n \n+    def requirements(self):\n+        self.requires(self.tested_reference_str)\n+\n     def build(self):\n         cmake = CMake(self)\n         cmake.configure()"
		},
		{
		  "sha": "0d20897301b68bdd7b7c0a6fe54ad74b0e86c7f9",
		  "filename": "recipes/cpp-peglib/1.x.x/test_v1_package/CMakeLists.txt",
		  "status": "modified",
		  "additions": 4,
		  "deletions": 7,
		  "changes": 11,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/08c0d0373f5008624a472098615d5de48d056c0a/recipes%2Fcpp-peglib%2F1.x.x%2Ftest_v1_package%2FCMakeLists.txt",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/08c0d0373f5008624a472098615d5de48d056c0a/recipes%2Fcpp-peglib%2F1.x.x%2Ftest_v1_package%2FCMakeLists.txt",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fcpp-peglib%2F1.x.x%2Ftest_v1_package%2FCMakeLists.txt?ref=08c0d0373f5008624a472098615d5de48d056c0a",
		  "patch": "@@ -1,11 +1,8 @@\n-cmake_minimum_required(VERSION 3.8)\n-project(test_package LANGUAGES CXX)\n+cmake_minimum_required(VERSION 3.1)\n+project(test_package)\n \n include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)\n conan_basic_setup(TARGETS)\n \n-find_package(cpp-peglib CONFIG REQUIRED)\n-\n-add_executable(${PROJECT_NAME} ../test_package/test_package.cpp)\n-target_link_libraries(${PROJECT_NAME} PRIVATE cpp-peglib::cpp-peglib)\n-target_compile_features(${PROJECT_NAME} PRIVATE cxx_std_17)\n+add_subdirectory(${CMAKE_CURRENT_SOURCE_DIR}/../test_package\n+                 ${CMAKE_CURRENT_BINARY_DIR}/test_package)"
		}
	  ]`)

	obtained, err := processChangedFiles(files)
	assert.Equal(t, err, nil)
	assert.Equal(t, &change{Recipe: "cpp-peglib", Change: EDIT, Weight: REGULAR}, obtained)
}

func TestProcessChangedFiles16144(t *testing.T) {
	// https://api.github.com/repos/conan-io/conan-center-index/pulls/16144/files
	files := parsePullRequestFilesJSON(t, `[
		{
		  "sha": "c8ad0901b061ac1977cb8f354a1c4d124474a548",
		  "filename": "recipes/re2/all/conanfile.py",
		  "status": "modified",
		  "additions": 5,
		  "deletions": 7,
		  "changes": 12,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/e2aa65c961d48d688dd5450811229eb1d62649ba/recipes%2Fre2%2Fall%2Fconanfile.py",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/e2aa65c961d48d688dd5450811229eb1d62649ba/recipes%2Fre2%2Fall%2Fconanfile.py",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fre2%2Fall%2Fconanfile.py?ref=e2aa65c961d48d688dd5450811229eb1d62649ba",
		  "patch": "@@ -4,17 +4,18 @@\n from conan.tools.files import copy, get, rmdir\n import os\n \n-required_conan_version = \">=1.53.0\"\n+required_conan_version = \">=1.54.0\"\n \n \n class Re2Conan(ConanFile):\n     name = \"re2\"\n     description = \"Fast, safe, thread-friendly regular expression library\"\n-    topics = (\"regex\")\n+    topics = (\"regex\",)\n     url = \"https://github.com/conan-io/conan-center-index\"\n     homepage = \"https://github.com/google/re2\"\n     license = \"BSD-3-Clause\"\n \n+    package_type = \"library\"\n     settings = \"os\", \"arch\", \"compiler\", \"build_type\"\n     options = {\n         \"shared\": [True, False],\n@@ -37,18 +38,15 @@ def layout(self):\n         cmake_layout(self, src_folder=\"src\")\n \n     def validate(self):\n-        if self.info.settings.compiler.get_safe(\"cppstd\"):\n+        if self.settings.compiler.get_safe(\"cppstd\"):\n             check_min_cppstd(self, 11)\n \n     def source(self):\n-        get(self, **self.conan_data[\"sources\"][self.version],\n-            destination=self.source_folder, strip_root=True)\n+        get(self, **self.conan_data[\"sources\"][self.version], strip_root=True)\n \n     def generate(self):\n         tc = CMakeToolchain(self)\n         tc.variables[\"RE2_BUILD_TESTING\"] = False\n-        # Honor BUILD_SHARED_LIBS from conan_toolchain (see https://github.com/conan-io/conan/issues/11840)\n-        tc.cache_variables[\"CMAKE_POLICY_DEFAULT_CMP0077\"] = \"NEW\"\n         tc.generate()\n \n     def build(self):"
		}
	  ]`)

	obtained, err := processChangedFiles(files)
	assert.Equal(t, err, nil)
	assert.Equal(t, &change{Recipe: "re2", Change: EDIT, Weight: SMALL}, obtained)
}
