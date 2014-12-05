// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DO NOT EDIT THIS FILE. GENERATED BY mkdoc.sh.
// Edit the documentation in other files and rerun mkdoc.sh to generate this one.

/*
Go is a tool for managing Go source code.

Usage:

	go command [arguments]

The commands are:

    build       compile packages and dependencies
    clean       remove object files
    env         print Go environment information
    fix         run go tool fix on packages
    fmt         run gofmt on package sources
    generate    generate Go files by processing source
    get         download and install packages and dependencies
    install     compile and install packages and dependencies
    list        list packages
    run         compile and run Go program
    test        test packages
    tool        run specified go tool
    version     print Go version
    vet         run go tool vet on packages

Use "go help [command]" for more information about a command.

Additional help topics:

    c           calling between Go and C
    filetype    file types
    gopath      GOPATH environment variable
    importpath  import path syntax
    packages    description of package lists
    testflag    description of testing flags
    testfunc    description of testing functions

Use "go help [topic]" for more information about that topic.


Compile packages and dependencies

Usage:

	go build [-o output] [-i] [build flags] [packages]

Build compiles the packages named by the import paths,
along with their dependencies, but it does not install the results.

If the arguments are a list of .go files, build treats them as a list
of source files specifying a single package.

When the command line specifies a single main package,
build writes the resulting executable to output.
Otherwise build compiles the packages but discards the results,
serving only as a check that the packages can be built.

The -o flag specifies the output file name. If not specified, the
output file name depends on the arguments and derives from the name
of the package, such as p.a for package p, unless p is 'main'. If
the package is main and file names are provided, the file name
derives from the first file name mentioned, such as f1 for 'go build
f1.go f2.go'; with no files provided ('go build'), the output file
name is the base name of the containing directory.

The -i flag installs the packages that are dependencies of the target.

The build flags are shared by the build, clean, get, install, list, run,
and test commands:

	-a
		force rebuilding of packages that are already up-to-date.
		In Go releases, does not apply to the standard library.
	-n
		print the commands but do not run them.
	-p n
		the number of builds that can be run in parallel.
		The default is the number of CPUs available.
	-race
		enable data race detection.
		Supported only on linux/amd64, freebsd/amd64, darwin/amd64 and windows/amd64.
	-v
		print the names of packages as they are compiled.
	-work
		print the name of the temporary work directory and
		do not delete it when exiting.
	-x
		print the commands.

	-ccflags 'arg list'
		arguments to pass on each 5c, 6c, or 8c compiler invocation.
	-compiler name
		name of compiler to use, as in runtime.Compiler (gccgo or gc).
	-gccgoflags 'arg list'
		arguments to pass on each gccgo compiler/linker invocation.
	-gcflags 'arg list'
		arguments to pass on each 5g, 6g, or 8g compiler invocation.
	-installsuffix suffix
		a suffix to use in the name of the package installation directory,
		in order to keep output separate from default builds.
		If using the -race flag, the install suffix is automatically set to race
		or, if set explicitly, has _race appended to it.
	-ldflags 'flag list'
		arguments to pass on each 5l, 6l, or 8l linker invocation.
	-tags 'tag list'
		a list of build tags to consider satisfied during the build.
		For more information about build tags, see the description of
		build constraints in the documentation for the go/build package.

The list flags accept a space-separated list of strings. To embed spaces
in an element in the list, surround it with either single or double quotes.

For more about specifying packages, see 'go help packages'.
For more about where packages and binaries are installed,
run 'go help gopath'.  For more about calling between Go and C/C++,
run 'go help c'.

See also: go install, go get, go clean.


Remove object files

Usage:

	go clean [-i] [-r] [-n] [-x] [build flags] [packages]

Clean removes object files from package source directories.
The go command builds most objects in a temporary directory,
so go clean is mainly concerned with object files left by other
tools or by manual invocations of go build.

Specifically, clean removes the following files from each of the
source directories corresponding to the import paths:

	_obj/            old object directory, left from Makefiles
	_test/           old test directory, left from Makefiles
	_testmain.go     old gotest file, left from Makefiles
	test.out         old test log, left from Makefiles
	build.out        old test log, left from Makefiles
	*.[568ao]        object files, left from Makefiles

	DIR(.exe)        from go build
	DIR.test(.exe)   from go test -c
	MAINFILE(.exe)   from go build MAINFILE.go
	*.so             from SWIG

In the list, DIR represents the final path element of the
directory, and MAINFILE is the base name of any Go source
file in the directory that is not included when building
the package.

The -i flag causes clean to remove the corresponding installed
archive or binary (what 'go install' would create).

The -n flag causes clean to print the remove commands it would execute,
but not run them.

The -r flag causes clean to be applied recursively to all the
dependencies of the packages named by the import paths.

The -x flag causes clean to print remove commands as it executes them.

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.


Print Go environment information

Usage:

	go env [var ...]

Env prints Go environment information.

By default env prints information as a shell script
(on Windows, a batch file).  If one or more variable
names is given as arguments,  env prints the value of
each named variable on its own line.


Run go tool fix on packages

Usage:

	go fix [packages]

Fix runs the Go fix command on the packages named by the import paths.

For more about fix, see 'godoc fix'.
For more about specifying packages, see 'go help packages'.

To run fix with specific options, run 'go tool fix'.

See also: go fmt, go vet.


Run gofmt on package sources

Usage:

	go fmt [-n] [-x] [packages]

Fmt runs the command 'gofmt -l -w' on the packages named
by the import paths.  It prints the names of the files that are modified.

For more about gofmt, see 'godoc gofmt'.
For more about specifying packages, see 'go help packages'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

To run gofmt with specific options, run gofmt itself.

See also: go fix, go vet.


Generate Go files by processing source

Usage:

	go generate [-run regexp] [file.go... | packages]

Generate runs commands described by directives within existing
files. Those commands can run any process but the intent is to
create or update Go source files, for instance by running yacc.

Go generate is never run automatically by go build, go get, go test,
and so on. It must be run explicitly.

Go generate scans the file for directives, which are lines of
the form,

	//go:generate command argument...

(note: no leading spaces and no space in "//go") where command
is the generator to be run, corresponding to an executable file
that can be run locally. It must either be in the shell path
(gofmt), a fully qualified path (/usr/you/bin/mytool), or a
command alias, described below.

Note that go generate does not parse the file, so lines that look
like directives in comments or multiline strings will be treated
as directives.

The arguments to the directive are space-separated tokens or
double-quoted strings passed to the generator as individual
arguments when it is run.

Quoted strings use Go syntax and are evaluated before execution; a
quoted string appears as a single argument to the generator.

Go generate sets several variables when it runs the generator:

	$GOARCH
		The execution architecture (arm, amd64, etc.)
	$GOOS
		The execution operating system (linux, windows, etc.)
	$GOFILE
		The base name of the file.
	$GOPACKAGE
		The name of the package of the file containing the directive.

Other than variable substitution and quoted-string evaluation, no
special processing such as "globbing" is performed on the command
line.

As a last step before running the command, any invocations of any
environment variables with alphanumeric names, such as $GOFILE or
$HOME, are expanded throughout the command line. The syntax for
variable expansion is $NAME on all operating systems.  Due to the
order of evaluation, variables are expanded even inside quoted
strings. If the variable NAME is not set, $NAME expands to the
empty string.

A directive of the form,

	//go:generate -command xxx args...

specifies, for the remainder of this source file only, that the
string xxx represents the command identified by the arguments. This
can be used to create aliases or to handle multiword generators.
For example,

	//go:generate -command yacc go tool yacc

specifies that the command "yacc" represents the generator
"go tool yacc".

Generate processes packages in the order given on the command line,
one at a time. If the command line lists .go files, they are treated
as a single package. Within a package, generate processes the
source files in a package in file name order, one at a time. Within
a source file, generate runs generators in the order they appear
in the file, one at a time.

If any generator returns an error exit status, "go generate" skips
all further processing for that package.

The generator is run in the package's source directory.

Go generate accepts one specific flag:

	-run=""
		if non-empty, specifies a regular expression to
		select directives whose command matches the expression.

It also accepts the standard build flags -v, -n, and -x.
The -v flag prints the names of packages and files as they are
processed.
The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

For more about specifying packages, see 'go help packages'.


Download and install packages and dependencies

Usage:

	go get [-d] [-f] [-fix] [-t] [-u] [build flags] [packages]

Get downloads and installs the packages named by the import paths,
along with their dependencies.

The -d flag instructs get to stop after downloading the packages; that is,
it instructs get not to install the packages.

The -f flag, valid only when -u is set, forces get -u not to verify that
each package has been checked out from the source control repository
implied by its import path. This can be useful if the source is a local fork
of the original.

The -fix flag instructs get to run the fix tool on the downloaded packages
before resolving dependencies or building the code.

The -t flag instructs get to also download the packages required to build
the tests for the specified packages.

The -u flag instructs get to use the network to update the named packages
and their dependencies.  By default, get uses the network to check out
missing packages but does not use it to look for updates to existing packages.

Get also accepts build flags to control the installation. See 'go help build'.

When checking out or updating a package, get looks for a branch or tag
that matches the locally installed version of Go. The most important
rule is that if the local installation is running version "go1", get
searches for a branch or tag named "go1". If no such version exists it
retrieves the most recent version of the package.

For more about specifying packages, see 'go help packages'.

For more about how 'go get' finds source code to
download, see 'go help importpath'.

See also: go build, go install, go clean.


Compile and install packages and dependencies

Usage:

	go install [build flags] [packages]

Install compiles and installs the packages named by the import paths,
along with their dependencies.

For more about the build flags, see 'go help build'.
For more about specifying packages, see 'go help packages'.

See also: go build, go get, go clean.


List packages

Usage:

	go list [-e] [-f format] [-json] [build flags] [packages]

List lists the packages named by the import paths, one per line.

The default output shows the package import path:

    code.google.com/p/google-api-go-client/books/v1
    code.google.com/p/goauth2/oauth
    code.google.com/p/sqlite

The -f flag specifies an alternate format for the list, using the
syntax of package template.  The default output is equivalent to -f
'{{.ImportPath}}'. The struct being passed to the template is:

    type Package struct {
        Dir           string // directory containing package sources
        ImportPath    string // import path of package in dir
        ImportComment string // path in import comment on package statement
        Name          string // package name
        Doc           string // package documentation string
        Target        string // install path
        Goroot        bool   // is this package in the Go root?
        Standard      bool   // is this package part of the standard Go library?
        Stale         bool   // would 'go install' do anything for this package?
        Root          string // Go root or Go path dir containing this package

        // Source files
        GoFiles        []string // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
        CgoFiles       []string // .go sources files that import "C"
        IgnoredGoFiles []string // .go sources ignored due to build constraints
        CFiles         []string // .c source files
        CXXFiles       []string // .cc, .cxx and .cpp source files
        MFiles         []string // .m source files
        HFiles         []string // .h, .hh, .hpp and .hxx source files
        SFiles         []string // .s source files
        SwigFiles      []string // .swig files
        SwigCXXFiles   []string // .swigcxx files
        SysoFiles      []string // .syso object files to add to archive

        // Cgo directives
        CgoCFLAGS    []string // cgo: flags for C compiler
        CgoCPPFLAGS  []string // cgo: flags for C preprocessor
        CgoCXXFLAGS  []string // cgo: flags for C++ compiler
        CgoLDFLAGS   []string // cgo: flags for linker
        CgoPkgConfig []string // cgo: pkg-config names

        // Dependency information
        Imports []string // import paths used by this package
        Deps    []string // all (recursively) imported dependencies

        // Error information
        Incomplete bool            // this package or a dependency has an error
        Error      *PackageError   // error loading package
        DepsErrors []*PackageError // errors loading dependencies

        TestGoFiles  []string // _test.go files in package
        TestImports  []string // imports from TestGoFiles
        XTestGoFiles []string // _test.go files outside package
        XTestImports []string // imports from XTestGoFiles
    }

The template function "join" calls strings.Join.

The template function "context" returns the build context, defined as:

	type Context struct {
		GOARCH        string   // target architecture
		GOOS          string   // target operating system
		GOROOT        string   // Go root
		GOPATH        string   // Go path
		CgoEnabled    bool     // whether cgo can be used
		UseAllFiles   bool     // use files regardless of +build lines, file names
		Compiler      string   // compiler to assume when computing target paths
		BuildTags     []string // build constraints to match in +build lines
		ReleaseTags   []string // releases the current release is compatible with
		InstallSuffix string   // suffix to use in the name of the install dir
	}

For more information about the meaning of these fields see the documentation
for the go/build package's Context type.

The -json flag causes the package data to be printed in JSON format
instead of using the template format.

The -e flag changes the handling of erroneous packages, those that
cannot be found or are malformed.  By default, the list command
prints an error to standard error for each erroneous package and
omits the packages from consideration during the usual printing.
With the -e flag, the list command never prints errors to standard
error and instead processes the erroneous packages with the usual
printing.  Erroneous packages will have a non-empty ImportPath and
a non-nil Error field; other information may or may not be missing
(zeroed).

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.


Compile and run Go program

Usage:

	go run [build flags] [-exec xprog] gofiles... [arguments...]

Run compiles and runs the main package comprising the named Go source files.
A Go source file is defined to be a file ending in a literal ".go" suffix.

By default, 'go run' runs the compiled binary directly: 'a.out arguments...'.
If the -exec flag is given, 'go run' invokes the binary using xprog: 'xprog a.out arguments...'.
If the -exec flag is not given, GOOS or GOARCH is different from the system
default, and a program named go_$GOOS_$GOARCH_exec can be found
on the current search path, 'go run' invokes the binary using that program,
for example 'go_nacl_386_exec a.out arguments...'. This allows execution of
cross-compiled programs when a simulator or other execution method is
available.

For more about build flags, see 'go help build'.

See also: go build.


Test packages

Usage:

	go test [-c] [-i] [build and test flags] [packages] [flags for test binary]

'Go test' automates testing the packages named by the import paths.
It prints a summary of the test results in the format:

	ok   archive/tar   0.011s
	FAIL archive/zip   0.022s
	ok   compress/gzip 0.033s
	...

followed by detailed output for each failed package.

'Go test' recompiles each package along with any files with names matching
the file pattern "*_test.go".
Files whose names begin with "_" (including "_test.go") or "." are ignored.
These additional files can contain test functions, benchmark functions, and
example functions.  See 'go help testfunc' for more.
Each listed package causes the execution of a separate test binary.

Test files that declare a package with the suffix "_test" will be compiled as a
separate package, and then linked and run with the main test binary.

By default, go test needs no arguments.  It compiles and tests the package
with source in the current directory, including tests, and runs the tests.

The package is built in a temporary directory so it does not interfere with the
non-test installation.

In addition to the build flags, the flags handled by 'go test' itself are:

	-c
		Compile the test binary to pkg.test but do not run it
		(where pkg is the last element of the package's import path).
		The file name can be changed with the -o flag.

	-exec xprog
	    Run the test binary using xprog. The behavior is the same as
	    in 'go run'. See 'go help run' for details.

	-i
	    Install packages that are dependencies of the test.
	    Do not run the test.

	-o file
		Compile the test binary to the named file.
		The test still runs (unless -c or -i is specified).


The test binary also accepts flags that control execution of the test; these
flags are also accessible by 'go test'.  See 'go help testflag' for details.

If the test binary needs any other flags, they should be presented after the
package names. The go tool treats as a flag the first argument that begins with
a minus sign that it does not recognize itself; that argument and all subsequent
arguments are passed as arguments to the test binary.

For more about build flags, see 'go help build'.
For more about specifying packages, see 'go help packages'.

See also: go build, go vet.


Run specified go tool

Usage:

	go tool [-n] command [args...]

Tool runs the go tool command identified by the arguments.
With no arguments it prints the list of known tools.

The -n flag causes tool to print the command that would be
executed but not execute it.

For more about each tool command, see 'go tool command -h'.


Print Go version

Usage:

	go version

Version prints the Go version, as reported by runtime.Version.


Run go tool vet on packages

Usage:

	go vet [-n] [-x] [packages]

Vet runs the Go vet command on the packages named by the import paths.

For more about vet, see 'godoc golang.org/x/tools/cmd/vet'.
For more about specifying packages, see 'go help packages'.

To run the vet tool with specific options, run 'go tool vet'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

See also: go fmt, go fix.


Calling between Go and C

There are two different ways to call between Go and C/C++ code.

The first is the cgo tool, which is part of the Go distribution.  For
information on how to use it see the cgo documentation (godoc cmd/cgo).

The second is the SWIG program, which is a general tool for
interfacing between languages.  For information on SWIG see
http://swig.org/.  When running go build, any file with a .swig
extension will be passed to SWIG.  Any file with a .swigcxx extension
will be passed to SWIG with the -c++ option.

When either cgo or SWIG is used, go build will pass any .c, .m, .s,
or .S files to the C compiler, and any .cc, .cpp, .cxx files to the C++
compiler.  The CC or CXX environment variables may be set to determine
the C or C++ compiler, respectively, to use.


File types

The go command examines the contents of a restricted set of files
in each directory. It identifies which files to examine based on
the extension of the file name. These extensions are:

	.go
		Go source files.
	.c, .h
		C source files.
		If the package uses cgo, these will be compiled with the
		OS-native compiler (typically gcc); otherwise they will be
		compiled with the Go-specific support compiler,
		5c, 6c, or 8c, etc. as appropriate.
	.cc, .cpp, .cxx, .hh, .hpp, .hxx
		C++ source files. Only useful with cgo or SWIG, and always
		compiled with the OS-native compiler.
	.m
		Objective-C source files. Only useful with cgo, and always
		compiled with the OS-native compiler.
	.s, .S
		Assembler source files.
		If the package uses cgo, these will be assembled with the
		OS-native assembler (typically gcc (sic)); otherwise they
		will be assembled with the Go-specific support assembler,
		5a, 6a, or 8a, etc., as appropriate.
	.swig, .swigcxx
		SWIG definition files.
	.syso
		System object files.

Files of each of these types except .syso may contain build
constraints, but the go command stops scanning for build constraints
at the first item in the file that is not a blank line or //-style
line comment.


GOPATH environment variable

The Go path is used to resolve import statements.
It is implemented by and documented in the go/build package.

The GOPATH environment variable lists places to look for Go code.
On Unix, the value is a colon-separated string.
On Windows, the value is a semicolon-separated string.
On Plan 9, the value is a list.

GOPATH must be set to get, build and install packages outside the
standard Go tree.

Each directory listed in GOPATH must have a prescribed structure:

The src/ directory holds source code.  The path below 'src'
determines the import path or executable name.

The pkg/ directory holds installed package objects.
As in the Go tree, each target operating system and
architecture pair has its own subdirectory of pkg
(pkg/GOOS_GOARCH).

If DIR is a directory listed in the GOPATH, a package with
source in DIR/src/foo/bar can be imported as "foo/bar" and
has its compiled form installed to "DIR/pkg/GOOS_GOARCH/foo/bar.a".

The bin/ directory holds compiled commands.
Each command is named for its source directory, but only
the final element, not the entire path.  That is, the
command with source in DIR/src/foo/quux is installed into
DIR/bin/quux, not DIR/bin/foo/quux.  The foo/ is stripped
so that you can add DIR/bin to your PATH to get at the
installed commands.  If the GOBIN environment variable is
set, commands are installed to the directory it names instead
of DIR/bin.

Here's an example directory layout:

    GOPATH=/home/user/gocode

    /home/user/gocode/
        src/
            foo/
                bar/               (go code in package bar)
                    x.go
                quux/              (go code in package main)
                    y.go
        bin/
            quux                   (installed command)
        pkg/
            linux_amd64/
                foo/
                    bar.a          (installed package object)

Go searches each directory listed in GOPATH to find source code,
but new packages are always downloaded into the first directory
in the list.


Import path syntax

An import path (see 'go help packages') denotes a package
stored in the local file system.  In general, an import path denotes
either a standard package (such as "unicode/utf8") or a package
found in one of the work spaces (see 'go help gopath').

Relative import paths

An import path beginning with ./ or ../ is called a relative path.
The toolchain supports relative import paths as a shortcut in two ways.

First, a relative path can be used as a shorthand on the command line.
If you are working in the directory containing the code imported as
"unicode" and want to run the tests for "unicode/utf8", you can type
"go test ./utf8" instead of needing to specify the full path.
Similarly, in the reverse situation, "go test .." will test "unicode" from
the "unicode/utf8" directory. Relative patterns are also allowed, like
"go test ./..." to test all subdirectories. See 'go help packages' for details
on the pattern syntax.

Second, if you are compiling a Go program not in a work space,
you can use a relative path in an import statement in that program
to refer to nearby code also not in a work space.
This makes it easy to experiment with small multipackage programs
outside of the usual work spaces, but such programs cannot be
installed with "go install" (there is no work space in which to install them),
so they are rebuilt from scratch each time they are built.
To avoid ambiguity, Go programs cannot use relative import paths
within a work space.

Remote import paths

Certain import paths also
describe how to obtain the source code for the package using
a revision control system.

A few common code hosting sites have special syntax:

	Bitbucket (Git, Mercurial)

		import "bitbucket.org/user/project"
		import "bitbucket.org/user/project/sub/directory"

	GitHub (Git)

		import "github.com/user/project"
		import "github.com/user/project/sub/directory"

	Google Code Project Hosting (Git, Mercurial, Subversion)

		import "code.google.com/p/project"
		import "code.google.com/p/project/sub/directory"

		import "code.google.com/p/project.subrepository"
		import "code.google.com/p/project.subrepository/sub/directory"

	Launchpad (Bazaar)

		import "launchpad.net/project"
		import "launchpad.net/project/series"
		import "launchpad.net/project/series/sub/directory"

		import "launchpad.net/~user/project/branch"
		import "launchpad.net/~user/project/branch/sub/directory"

	IBM DevOps Services (Git)

		import "hub.jazz.net/git/user/project"
		import "hub.jazz.net/git/user/project/sub/directory"

For code hosted on other servers, import paths may either be qualified
with the version control type, or the go tool can dynamically fetch
the import path over https/http and discover where the code resides
from a <meta> tag in the HTML.

To declare the code location, an import path of the form

	repository.vcs/path

specifies the given repository, with or without the .vcs suffix,
using the named version control system, and then the path inside
that repository.  The supported version control systems are:

	Bazaar      .bzr
	Git         .git
	Mercurial   .hg
	Subversion  .svn

For example,

	import "example.org/user/foo.hg"

denotes the root directory of the Mercurial repository at
example.org/user/foo or foo.hg, and

	import "example.org/repo.git/foo/bar"

denotes the foo/bar directory of the Git repository at
example.org/repo or repo.git.

When a version control system supports multiple protocols,
each is tried in turn when downloading.  For example, a Git
download tries git://, then https://, then http://.

If the import path is not a known code hosting site and also lacks a
version control qualifier, the go tool attempts to fetch the import
over https/http and looks for a <meta> tag in the document's HTML
<head>.

The meta tag has the form:

	<meta name="go-import" content="import-prefix vcs repo-root">

The import-prefix is the import path corresponding to the repository
root. It must be a prefix or an exact match of the package being
fetched with "go get". If it's not an exact match, another http
request is made at the prefix to verify the <meta> tags match.

The vcs is one of "git", "hg", "svn", etc,

The repo-root is the root of the version control system
containing a scheme and not containing a .vcs qualifier.

For example,

	import "example.org/pkg/foo"

will result in the following request(s):

	https://example.org/pkg/foo?go-get=1 (preferred)
	http://example.org/pkg/foo?go-get=1  (fallback)

If that page contains the meta tag

	<meta name="go-import" content="example.org git https://code.org/r/p/exproj">

the go tool will verify that https://example.org/?go-get=1 contains the
same meta tag and then git clone https://code.org/r/p/exproj into
GOPATH/src/example.org.

New downloaded packages are written to the first directory
listed in the GOPATH environment variable (see 'go help gopath').

The go command attempts to download the version of the
package appropriate for the Go release being used.
Run 'go help install' for more.


Description of package lists

Many commands apply to a set of packages:

	go action [packages]

Usually, [packages] is a list of import paths.

An import path that is a rooted path or that begins with
a . or .. element is interpreted as a file system path and
denotes the package in that directory.

Otherwise, the import path P denotes the package found in
the directory DIR/src/P for some DIR listed in the GOPATH
environment variable (see 'go help gopath').

If no import paths are given, the action applies to the
package in the current directory.

There are three reserved names for paths that should not be used
for packages to be built with the go tool:

- "main" denotes the top-level package in a stand-alone executable.

- "all" expands to all package directories found in all the GOPATH
trees. For example, 'go list all' lists all the packages on the local
system.

- "std" is like all but expands to just the packages in the standard
Go library.

An import path is a pattern if it includes one or more "..." wildcards,
each of which can match any string, including the empty string and
strings containing slashes.  Such a pattern expands to all package
directories found in the GOPATH trees with names matching the
patterns.  As a special case, x/... matches x as well as x's subdirectories.
For example, net/... expands to net and packages in its subdirectories.

An import path can also name a package to be downloaded from
a remote repository.  Run 'go help importpath' for details.

Every package in a program must have a unique import path.
By convention, this is arranged by starting each path with a
unique prefix that belongs to you.  For example, paths used
internally at Google all begin with 'google', and paths
denoting remote repositories begin with the path to the code,
such as 'code.google.com/p/project'.

As a special case, if the package list is a list of .go files from a
single directory, the command is applied to a single synthesized
package made up of exactly those files, ignoring any build constraints
in those files and ignoring any other files in the directory.

Directory and file names that begin with "." or "_" are ignored
by the go tool, as are directories named "testdata".


Description of testing flags

The 'go test' command takes both flags that apply to 'go test' itself
and flags that apply to the resulting test binary.

Several of the flags control profiling and write an execution profile
suitable for "go tool pprof"; run "go tool pprof help" for more
information.  The --alloc_space, --alloc_objects, and --show_bytes
options of pprof control how the information is presented.

The following flags are recognized by the 'go test' command and
control the execution of any test:

	-bench regexp
	    Run benchmarks matching the regular expression.
	    By default, no benchmarks run. To run all benchmarks,
	    use '-bench .' or '-bench=.'.

	-benchmem
	    Print memory allocation statistics for benchmarks.

	-benchtime t
	    Run enough iterations of each benchmark to take t, specified
	    as a time.Duration (for example, -benchtime 1h30s).
	    The default is 1 second (1s).

	-blockprofile block.out
	    Write a goroutine blocking profile to the specified file
	    when all tests are complete.
	    Writes test binary as -c would.

	-blockprofilerate n
	    Control the detail provided in goroutine blocking profiles by
	    calling runtime.SetBlockProfileRate with n.
	    See 'godoc runtime SetBlockProfileRate'.
	    The profiler aims to sample, on average, one blocking event every
	    n nanoseconds the program spends blocked.  By default,
	    if -test.blockprofile is set without this flag, all blocking events
	    are recorded, equivalent to -test.blockprofilerate=1.

	-cover
	    Enable coverage analysis.

	-covermode set,count,atomic
	    Set the mode for coverage analysis for the package[s]
	    being tested. The default is "set" unless -race is enabled,
	    in which case it is "atomic".
	    The values:
		set: bool: does this statement run?
		count: int: how many times does this statement run?
		atomic: int: count, but correct in multithreaded tests;
			significantly more expensive.
	    Sets -cover.

	-coverpkg pkg1,pkg2,pkg3
	    Apply coverage analysis in each test to the given list of packages.
	    The default is for each test to analyze only the package being tested.
	    Packages are specified as import paths.
	    Sets -cover.

	-coverprofile cover.out
	    Write a coverage profile to the file after all tests have passed.
	    Sets -cover.

	-cpu 1,2,4
	    Specify a list of GOMAXPROCS values for which the tests or
	    benchmarks should be executed.  The default is the current value
	    of GOMAXPROCS.

	-cpuprofile cpu.out
	    Write a CPU profile to the specified file before exiting.
	    Writes test binary as -c would.

	-memprofile mem.out
	    Write a memory profile to the file after all tests have passed.
	    Writes test binary as -c would.

	-memprofilerate n
	    Enable more precise (and expensive) memory profiles by setting
	    runtime.MemProfileRate.  See 'godoc runtime MemProfileRate'.
	    To profile all memory allocations, use -test.memprofilerate=1
	    and pass --alloc_space flag to the pprof tool.

	-outputdir directory
	    Place output files from profiling in the specified directory,
	    by default the directory in which "go test" is running.

	-parallel n
	    Allow parallel execution of test functions that call t.Parallel.
	    The value of this flag is the maximum number of tests to run
	    simultaneously; by default, it is set to the value of GOMAXPROCS.

	-run regexp
	    Run only those tests and examples matching the regular
	    expression.

	-short
	    Tell long-running tests to shorten their run time.
	    It is off by default but set during all.bash so that installing
	    the Go tree can run a sanity check but not spend time running
	    exhaustive tests.

	-timeout t
	    If a test runs longer than t, panic.

	-v
	    Verbose output: log all tests as they are run. Also print all
	    text from Log and Logf calls even if the test succeeds.

The test binary, called pkg.test where pkg is the name of the
directory containing the package sources, can be invoked directly
after building it with 'go test -c'. When invoking the test binary
directly, each of the standard flag names must be prefixed with 'test.',
as in -test.run=TestMyFunc or -test.v.

When running 'go test', flags not listed above are passed through
unaltered. For instance, the command

	go test -x -v -cpuprofile=prof.out -dir=testdata -update

will compile the test binary and then run it as

	pkg.test -test.v -test.cpuprofile=prof.out -dir=testdata -update

The test flags that generate profiles (other than for coverage) also
leave the test binary in pkg.test for use when analyzing the profiles.

Flags not recognized by 'go test' must be placed after any specified packages.


Description of testing functions

The 'go test' command expects to find test, benchmark, and example functions
in the "*_test.go" files corresponding to the package under test.

A test function is one named TestXXX (where XXX is any alphanumeric string
not starting with a lower case letter) and should have the signature,

	func TestXXX(t *testing.T) { ... }

A benchmark function is one named BenchmarkXXX and should have the signature,

	func BenchmarkXXX(b *testing.B) { ... }

An example function is similar to a test function but, instead of using
*testing.T to report success or failure, prints output to os.Stdout.
That output is compared against the function's "Output:" comment, which
must be the last comment in the function body (see example below). An
example with no such comment, or with no text after "Output:" is compiled
but not executed.

Godoc displays the body of ExampleXXX to demonstrate the use
of the function, constant, or variable XXX.  An example of a method M with
receiver type T or *T is named ExampleT_M.  There may be multiple examples
for a given function, constant, or variable, distinguished by a trailing _xxx,
where xxx is a suffix not beginning with an upper case letter.

Here is an example of an example:

	func ExamplePrintln() {
		Println("The output of\nthis example.")
		// Output: The output of
		// this example.
	}

The entire test file is presented as the example when it contains a single
example function, at least one other function, type, variable, or constant
declaration, and no test or benchmark functions.

See the documentation of the testing package for more information.


*/
package main
