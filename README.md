# Sudo Shell

Sudo Shell is a wrapper to run a login shell with `sudo` for the purpose of audit logging. 


## Purpose

The `sudo` command provides built-in session logging. Combined with [`sudoreplay`](https://www.sudo.ws/man/1.8.13/sudoreplay.man.html), there's an easy way to review system logs on a [bastion](https://github.com/cloudposse/bastion/) host. This wrapper can be used as a system login shell. By forcing the user login shell to call `sudo`, we can enforce audit logs for all users by default.

Another common pattern is to use the `ForceCommand` directive with `sshd_config` in OpenSSH. This is dangerous because if the [user can easily bypass it](http://serverfault.com/a/639814). 


## Usage

1. Enable `sudo` logging.

`/etc/sudoers.d/audit-logs`

```
Defaults log_output
Defaults!/usr/bin/sudoreplay !log_output
Defaults!/sbin/reboot !log_output
```

2. Add this command to `/etc/shells`

`/etc/shells`:

```
/usr/bin/sudosh
```

Tip: to prevent users from using other shells to login, remove them from `/etc/shells`.


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

Then set the user's shell to `/usr/bin/sudosh.zsh` and add the shell to `/etc/shells.

