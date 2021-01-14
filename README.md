# Gomail

**Requirements**

- Go
- Gmail account

**Installation**

Download and install gomail

```bash
go get github.com/bbriano/gomail
```

setup environment variables

```bash
export GOMAIL_USER='yourname@example.com'
```

Setup template file (Must have this file in home path!)

```bash
echo 'From: <your_email>
Subject: <subject>


Kind regards,
<yours_truly>' > ~/.gomail_template
```

**Usage**

Mail someone. Will use default editor

```bash
gomail someone@example.com
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
