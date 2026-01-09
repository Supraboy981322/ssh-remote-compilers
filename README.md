# Remote compilation over SSH

## Why?

For some reason, binaries that I compile on my PC (x86 NixOS) don't run on any other x86 Linux machine, but binaries compiled on other machines run fine on both it and any other x86 Linux machine. So in order to have portable binaries, I must compile on a separate machine. I was doing it semi-manually over SSH, but I got tired of writing the following command (for example)
```sh
cd ..; tar -cf - src/* | ssh super@dev 'set -e; mkdir -p /tmp/foo; cd /tmp/foo; tar -xf -; cd src; go build -o out.bin .;cat out.bin; cd /tmp; rm -r foo' > foo; cd src
```
This project is a server and client that does this automatically

## Usage

>[!NOTE]
>The client isn't written yet, but you can still use it with the SSH command

- The user you connect as is the compiler the server will use.

  For example, `go@your-server` will use the Go compiler.

- The final binary is dumped to the stdout of your connection, so make sure to pipe it to a file.

- Whatever commands passed when connecting to the server are passed as arguments to the compiler.
  
  For example, in order to compile a Go project from directory, you must pipe a tarball of the project directory to the SSH connection and pass `.` to the server as a command.
(replace `your-server` with the address for your SSH server and `foo` with your binary name)
```sh
tar -cf - . | ssh go@your-server -p 7845 . > foo
```

- Also, just as a warning, if you pass an argument to the server that sets the output file, you will recieve an empty binary. This is because the server sets the output file name to a random string internally, and your arguments are appended after the (for example) `-o er89jht84wt.bin` argument.

- Finally, It should be noted that, until the client is written, you'll still have to make it executable, you can just add `sudo chmod a+x foo` to the end of the chain for this.

Once the client is written, usage is likely going to be more like this:
```sh
rC go -o foo .
```
