# caddy-save
caddy handler to save request body to specified file

# Usage
## Self-updating Caddyfile with [caddy-exec](https://github.com/abiosoft/caddy-exec)
```caddy
route /caddy {
	basic_auth {
		Deployment <hashed password>
	}
	save Caddyfile
	exec caddy reload
}
```
