// +build debug

package assets

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
	"path/filepath"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// assets_client_page_html reads file data from disk. It returns an error on failure.
func assets_client_page_html() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/page.html"
	name := "assets/client/page.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_static_css_bootstrap_min_css reads file data from disk. It returns an error on failure.
func assets_client_static_css_bootstrap_min_css() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/static/css/bootstrap.min.css"
	name := "assets/client/static/css/bootstrap.min.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_static_css_highlight_min_css reads file data from disk. It returns an error on failure.
func assets_client_static_css_highlight_min_css() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/static/css/highlight.min.css"
	name := "assets/client/static/css/highlight.min.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_static_img_glyphicons_halflings_png reads file data from disk. It returns an error on failure.
func assets_client_static_img_glyphicons_halflings_png() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/static/img/glyphicons-halflings.png"
	name := "assets/client/static/img/glyphicons-halflings.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_static_js_angular_sanitize_min_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_angular_sanitize_min_js() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/static/js/angular-sanitize.min.js"
	name := "assets/client/static/js/angular-sanitize.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_static_js_angular_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_angular_js() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/static/js/angular.js"
	name := "assets/client/static/js/angular.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_static_js_base64_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_base64_js() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/static/js/base64.js"
	name := "assets/client/static/js/base64.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_static_js_highlight_min_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_highlight_min_js() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/static/js/highlight.min.js"
	name := "assets/client/static/js/highlight.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_static_js_jquery_1_9_1_min_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_jquery_1_9_1_min_js() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/static/js/jquery-1.9.1.min.js"
	name := "assets/client/static/js/jquery-1.9.1.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_static_js_jquery_timeago_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_jquery_timeago_js() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/static/js/jquery.timeago.js"
	name := "assets/client/static/js/jquery.timeago.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_static_js_ngrok_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_ngrok_js() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/static/js/ngrok.js"
	name := "assets/client/static/js/ngrok.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_static_js_vkbeautify_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_vkbeautify_js() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/static/js/vkbeautify.js"
	name := "assets/client/static/js/vkbeautify.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_tls_ngrokroot_crt reads file data from disk. It returns an error on failure.
func assets_client_tls_ngrokroot_crt() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/tls/ngrokroot.crt"
	name := "assets/client/tls/ngrokroot.crt"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_client_tls_snakeoilca_crt reads file data from disk. It returns an error on failure.
func assets_client_tls_snakeoilca_crt() (*asset, error) {
	path := "/home/way/gocode/src/ngrok/assets/client/tls/snakeoilca.crt"
	name := "assets/client/tls/snakeoilca.crt"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/client/page.html": assets_client_page_html,
	"assets/client/static/css/bootstrap.min.css": assets_client_static_css_bootstrap_min_css,
	"assets/client/static/css/highlight.min.css": assets_client_static_css_highlight_min_css,
	"assets/client/static/img/glyphicons-halflings.png": assets_client_static_img_glyphicons_halflings_png,
	"assets/client/static/js/angular-sanitize.min.js": assets_client_static_js_angular_sanitize_min_js,
	"assets/client/static/js/angular.js": assets_client_static_js_angular_js,
	"assets/client/static/js/base64.js": assets_client_static_js_base64_js,
	"assets/client/static/js/highlight.min.js": assets_client_static_js_highlight_min_js,
	"assets/client/static/js/jquery-1.9.1.min.js": assets_client_static_js_jquery_1_9_1_min_js,
	"assets/client/static/js/jquery.timeago.js": assets_client_static_js_jquery_timeago_js,
	"assets/client/static/js/ngrok.js": assets_client_static_js_ngrok_js,
	"assets/client/static/js/vkbeautify.js": assets_client_static_js_vkbeautify_js,
	"assets/client/tls/ngrokroot.crt": assets_client_tls_ngrokroot_crt,
	"assets/client/tls/snakeoilca.crt": assets_client_tls_snakeoilca_crt,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"client": &_bintree_t{nil, map[string]*_bintree_t{
			"page.html": &_bintree_t{assets_client_page_html, map[string]*_bintree_t{
			}},
			"static": &_bintree_t{nil, map[string]*_bintree_t{
				"css": &_bintree_t{nil, map[string]*_bintree_t{
					"bootstrap.min.css": &_bintree_t{assets_client_static_css_bootstrap_min_css, map[string]*_bintree_t{
					}},
					"highlight.min.css": &_bintree_t{assets_client_static_css_highlight_min_css, map[string]*_bintree_t{
					}},
				}},
				"img": &_bintree_t{nil, map[string]*_bintree_t{
					"glyphicons-halflings.png": &_bintree_t{assets_client_static_img_glyphicons_halflings_png, map[string]*_bintree_t{
					}},
				}},
				"js": &_bintree_t{nil, map[string]*_bintree_t{
					"angular-sanitize.min.js": &_bintree_t{assets_client_static_js_angular_sanitize_min_js, map[string]*_bintree_t{
					}},
					"angular.js": &_bintree_t{assets_client_static_js_angular_js, map[string]*_bintree_t{
					}},
					"base64.js": &_bintree_t{assets_client_static_js_base64_js, map[string]*_bintree_t{
					}},
					"highlight.min.js": &_bintree_t{assets_client_static_js_highlight_min_js, map[string]*_bintree_t{
					}},
					"jquery-1.9.1.min.js": &_bintree_t{assets_client_static_js_jquery_1_9_1_min_js, map[string]*_bintree_t{
					}},
					"jquery.timeago.js": &_bintree_t{assets_client_static_js_jquery_timeago_js, map[string]*_bintree_t{
					}},
					"ngrok.js": &_bintree_t{assets_client_static_js_ngrok_js, map[string]*_bintree_t{
					}},
					"vkbeautify.js": &_bintree_t{assets_client_static_js_vkbeautify_js, map[string]*_bintree_t{
					}},
				}},
			}},
			"tls": &_bintree_t{nil, map[string]*_bintree_t{
				"ngrokroot.crt": &_bintree_t{assets_client_tls_ngrokroot_crt, map[string]*_bintree_t{
				}},
				"snakeoilca.crt": &_bintree_t{assets_client_tls_snakeoilca_crt, map[string]*_bintree_t{
				}},
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

