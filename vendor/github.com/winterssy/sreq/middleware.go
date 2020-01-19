package sreq

// SetDefaultHost is a before request hook set client level host,
// can be overwrite by request level option.
func SetDefaultHost(host string) BeforeRequestHook {
	return func(req *Request) error {
		if req.Host == "" {
			req.SetHost(host)
		}
		return nil
	}
}

// SetDefaultHeaders is a before request hook to set client level headers,
// can be overwrite by request level option.
func SetDefaultHeaders(headers Headers) BeforeRequestHook {
	return func(req *Request) error {
		req.Headers.Merge(headers)
		return nil
	}
}

// SetDefaultContentType is a before request hook to set client level Content-Type header value,
// can be overwrite by request level option.
func SetDefaultContentType(contentType string) BeforeRequestHook {
	return func(req *Request) error {
		req.Headers.SetDefault("Content-Type", contentType)
		return nil
	}
}

// SetDefaultUserAgent is a before request hook to set client level User-Agent header value,
// can be overwrite by request level option.
func SetDefaultUserAgent(userAgent string) BeforeRequestHook {
	return func(req *Request) error {
		req.Headers.SetDefault("User-Agent", userAgent)
		return nil
	}
}

// SetDefaultReferer is a before request hook to set client level Referer header value,
// can be overwrite by request level option.
func SetDefaultReferer(referer string) BeforeRequestHook {
	return func(req *Request) error {
		req.Headers.SetDefault("Referer", referer)
		return nil
	}
}

// SetDefaultQuery is a before request hook to set client level query parameters,
// can be overwrite by request level option.
func SetDefaultQuery(query Params) BeforeRequestHook {
	return func(req *Request) error {
		req.Query.Merge(query)
		return nil
	}
}

// SetDefaultForm is a before request hook to set client level form payload,
// can be overwrite by request level option.
func SetDefaultForm(form Form) BeforeRequestHook {
	return func(req *Request) error {
		req.Form.Merge(form)
		return nil
	}
}

// SetDefaultCookies is a before request hook to set client level cookies,
// can be overwrite by request level option.
func SetDefaultCookies(cookies Cookies) BeforeRequestHook {
	return func(req *Request) error {
		req.Cookies.Merge(cookies)
		return nil
	}
}

// SetDefaultBasicAuth is a before request hook to set client level basic authentication,
// can be overwrite by request level option.
func SetDefaultBasicAuth(username string, password string) BeforeRequestHook {
	return func(req *Request) error {
		req.Headers.SetDefault("Authorization", "Basic "+basicAuth(username, password))
		return nil
	}
}

// SetDefaultBearerToken is a before request hook to set client level bearer token,
// can be overwrite by request level option.
func SetDefaultBearerToken(token string) BeforeRequestHook {
	return func(req *Request) error {
		req.Headers.SetDefault("Authorization", "Bearer "+token)
		return nil
	}
}

// SetDefaultRetry is a before request hook to set client level retry policy,
// can be overwrite by request level option.
func SetDefaultRetry(retry *Retry) BeforeRequestHook {
	return func(req *Request) error {
		req.Retry = req.Retry.Merge(retry)
		return nil
	}
}