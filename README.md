
# Sudo Shell

Sudo Shell is a wrapper to run a login shell with `sudo` for the purpose of audit logging. 

[![Build Status](https://travis-ci.org/cloudposse/sudosh.svg?branch=master)](https://travis-ci.org/cloudposse/sudosh)
[![GitHub Stars](https://img.shields.io/github/stars/cloudposse/sudosh.svg)](https://github.com/cloudposse/sudosh/stargazers) 
[![GitHub Issues](https://img.shields.io/github/issues/cloudposse/sudosh.svg)](https://github.com/cloudposse/sudosh/issues)
[![Average time to resolve an issue](http://isitmaintained.com/badge/resolution/cloudposse/sudosh.svg)](http://isitmaintained.com/project/cloudposse/sudosh "Average time to resolve an issue")
[![Percentage of issues still open](http://isitmaintained.com/badge/open/cloudposse/sudosh.svg)](http://isitmaintained.com/project/cloudposse/sudosh "Percentage of issues still open")
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg)](https://github.com/cloudposse/sudosh/pulls)
[![License](https://img.shields.io/badge/license-APACHE%202.0%20-brightgreen.svg)](https://github.com/cloudposse/sudosh/blob/master/LICENSE)


## Purpose

The `sudo` command provides built-in session logging. Combined with [`sudoreplay`](https://www.sudo.ws/man/1.8.13/sudoreplay.man.html), there's an easy way to review system logs on a [bastion](https://github.com/cloudposse/bastion/) host. This wrapper can be used as a system login shell. By forcing the user login shell to call `sudo`, we can enforce audit logs for all users by default.

Another common pattern is to use the `ForceCommand` directive with `sshd_config` in OpenSSH. This is dangerous because if the [user can easily bypass it](http://serverfault.com/a/639814). 


## Usage

1. Enable `sudo` logging. Edit `/etc/sudoers.d/audit-logs`:

```
Defaults log_output
Defaults!/usr/bin/sudoreplay !log_output
Defaults!/sbin/reboot !log_output
```

2. Add this command to `/etc/shells`:

```
/usr/bin/sudosh
```

**Tip**: to prevent users from using other shells to login, remove them from `/etc/shells`.


3. Update the user `foobar` to use the `sudosh` shell.

```
chsh -s /usr/bin/sudosh foobar
```


## Other Tricks

If you want to change the default shell from `bash` to something else (e.g. `zsh`), you can symlink `sudosh` to a different name. 

To change the default shell to `zsh`, you could do:

```
ln -s /usr/bin/sudosh /usr/bin/sudosh.zsh
```

Then set the user's shell to `/usr/bin/sudosh.zsh` and add the shell to `/etc/shells`.

