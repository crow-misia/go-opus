### Build Environment install

download and install `msys2-x86_64-yyyymmdd.exe` from  http://www.msys2.org/

```
pacman -Syuu
pacman -S base-devel
pacman -S mingw-w64-x86_64-toolchain
pacman -S mingw64/mingw-w64-x86_64-cmake
```

### opus install

```
git clone https://github.com/xiph/opus.git
cd opus
mkdir build
cd build
cmake -G"MSYS Makefiles" -D OPUS_ENABLE_FLOAT_API=1 -D OPUS_STACK_PROTECTOR=0 -D CMAKE_INSTALL_PREFIX=/mingw64 ..
make
make install
```

### Go build

```
go build examples/opus_demo/main.go
```