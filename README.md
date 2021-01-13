# Gomail

**Requirements**

- Go

**Installation**

```bash
go get github.com/bbriano/gomail
echo 'Hello,\n' > ~/.gomail_template
```

currently only works with google's mail server

setup environment variables

```bash
export GOMAIL_USER='yourname@example.com'
export GOMAIL_PASS='Password123'
```

**Usage**

Mail someone. Will use default editor

```bash
gomail someone@example.com
```

Mail to self

```bash
gomail
```

Will read from STDIN if its not empty

```bash
echo 'hello, world' | gomail someone@example.com
```

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
