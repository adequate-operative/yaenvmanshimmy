### Does thing like

    Reads `<executableName>.shimmy` file when launched and uses as config.

    Launchs a process based on arguments: system/shellexec's line in file.

    Lines startsing with `@ENV<some env var>` replaced with value of enviroment var.

    shimmy config files have obvious security implications e.g. if edited maliciously

### usage

    Build `cmd/shimmylaunch`

    Rename `shimmylaunch.exe` to `something_sensible.exe`

    Put a file named `something_sensible.exe.yaenvshimmy` in the same directory as `something_sensible.exe`

    `.yaenvshimmy` file lines correspond to launch arguments

    First list should be the full path of target/proxied executable

    Line beginging with `@ENV` replaced with env var name following `@ENV`. e.g. `@ENVPATH` ==becomes=> contents of `$PATH`

    Lines begining with `@PWD` are replaced with the working directory that shimmy's env was set to by calling proccess

### TODO

- Optionally unpack config from envvar e.g. `$SHIMMYPACK=Base85(encoded config file)`.
  Unpacking this variable should remove it from the envars (unless it is subsiquently replaced in the config)
  To Support chaining... why?

  - If use case includes launched exe calling other shimms. Not hard to impliment, but seems like a nightmare to use.

- Fancy Support for `$PATH` ? e.g. replace

- Support Changing env sent to launched executable: replace, remove, add

- Support Changing working directory

- generate executable with from config file

- Alternitively configure build with hardcoded password (HASH)?, config file requires matching HMAC. Config file would need a script to stamp after each edit.
