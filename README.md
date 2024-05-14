About w32
==========

# w32 Fork

**This project is a fork of the original [w32](https://github.com/AllenDang/w32) project, created for self-maintenance, development, and for use in internal projects.**

[![Go Report Card](https://goreportcard.com/badge/github.com/j4ackwolf/w32)](https://goreportcard.com/report/github.com/j4ackwolf/w32)

w32 is a wrapper of windows apis for the Go Programming Language.

It wraps win32 apis to "Go style" to make them easier to use.

Setup
=====

1. Make sure you have a working Go installation and build environment, 
   see this go-nuts post for details:
   http://groups.google.com/group/golang-nuts/msg/5c87630a84f4fd0c
   
   Updated versions of the Windows Go build are available here:
   http://code.google.com/p/gomingw/downloads/list
   
2. Create a "gopath" directory if you do not have one yet and set the
   GOPATH variable accordingly. For example:
   mkdir -p go-externals/src
   export GOPATH=${PWD}/go-externals

3. go get github.com/j4ackwolf/w32

4. go install github.com/j4ackwolf/w32...

Contribute
==========

Contributions in form of design, code, documentation, bug reporting or other
ways you see fit are very welcome.

## Thank You!

We would like to express our sincere gratitude to the original author, [Allen Dang](https://github.com/AllenDang), for creating the original w32 project and making it available for the community. Without their work, this fork would not have been possible.

**Thank you for your support and contributions!**

## Original Project

The original w32 project can be found [here](https://github.com/AllenDang/w32). Please consider checking out the original project for additional context and information.

## License

This project is licensed under the [MIT License](LICENSE).


