<h1 align="center">Discard Rest APIs</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Made%20with-%E2%9D%A4-red" alt="Made with Love">
  <img src="https://img.shields.io/badge/Open%20Source-%F0%9F%8C%90-blue" alt="Open Source">
</p>

<p align="center">
  A hub of RESTful APIs for developers, providing 500+ powerful endpoints across multiple categories.<br>
  From downloaders and AI tools to image processing, games, and converters - everything you need to elevate your applications to new heights.
</p>

## Features

- **Developer-Friendly** - Intuitive endpoints with clear documentation
- **Extensive Collection** - 500+ ready-to-use REST APIs across 30+ categories
- **Affordable Pricing** - Competitive rates with many free endpoints
- **Blazing Fast** - Optimized for speed and reliability
- **Seamless Integration** - Works with any programming language or framework
- **Complete Documentation** - Detailed guides and examples for every API

## API Categories

Discard offers APIs across numerous categories including:

- **Islamic** - Quran, Surah, Hadith and more
- **News** - Latest news endpoints of many platforms
- **Downloader** - Download content from many platforms
- **AI Tools** - Many ai models and llms
- **Image Processing** - Filters, effects, and more
- **Stalking** - Social media data extraction
- **Searching** - Advanced search apis across platforms
- **Converters** - Format and convert your data
- **And many more!**

Explore our complete catalog at [discardapi.dpdns.org](https://discardapi.dpdns.org)

## Getting Started with Discard Rest APIs

<details>
<summary><b>Click to expand</b></summary>

**Welcome to Discard Rest APIs, your one-stop solution for seamless API integrations! Our extensive collection of APIs is designed for developers building apps, businesses enhancing services, or tech enthusiasts experimenting with new ideas.**

### Step 1: Sign Up & Get Your API Key

- Create an account to access our API dashboard. Signing up is quick and easy, providing instant access to hundreds of powerful APIs.

### Step 2: Choose an API

- Browse our comprehensive API library and select the API that fits your needs. Each API includes detailed documentation with endpoints, parameters, and response formats.

### Step 3: Make Your First API Call

- With your API key in hand, you're ready to start! All our APIs follow REST principles and are designed for simple, intuitive integration.

### Step 4: Integrate the API

- Easily incorporate our APIs into your existing systems using the provided code examples for popular languages like JavaScript, Python, PHP and more.

### Step 5: Upgrade for More Features

- For extensive usage and advanced features, upgrade to a PRO or VIP plan offering higher limits, faster response times, and premium feature access.
</details>

---


## 🚀 Quick Start  

Get started in seconds. Add your **API key** to any endpoint and receive structured JSON responses with comprehensive error handling.    

---

### ⚙️ Request Format  
```http
GET https://discardapi.dpdns.org/api/endpoint?apikey=YOUR_KEY
```

```json

{
  "status": true,
  "creator": "Qasim Ali 🩷",
  "result": { ... }
}
```

### 🎴 Vignette Image Upload (🟣POST)

**Endpoint**
```nginx
POST https://discardapi.dpdns.org/api/image/vignette
```

**🔹Usage cURL**
```bash
curl -X POST "https://discardapi.dpdns.org/api/image/vignette" \
  -F "apikey=your_api_key" \
  -F "file=@input.jpg" \
  -F "intensity=0.8" \
  -F "shape=circle" \
  -F "color=#000000" \
  --output output.jpg

```
**🔹Usage JavaScript**
```js

import axios from "axios";
import fs from "fs";
import FormData from "form-data";

const form = new FormData();
form.append("apikey", "your_api_key");
form.append("file", fs.createReadStream("input.jpg"));
form.append("intensity", "0.8");
form.append("shape", "circle");
form.append("color", "#000000");

axios.post("https://discardapi.dpdns.org/api/image/vignette", form, {
  headers: form.getHeaders(),
  responseType: "stream"
})
.then(res => {
  const writer = fs.createWriteStream("output.jpg");
  res.data.pipe(writer);
})
.catch(err => console.error(err.response?.data || err.message));
```
**Usage Python**
```py
import requests

url = "https://discardapi.dpdns.org/api/image/vignette"
files = {"file": open("input.jpg", "rb")}
data = {
    "apikey": "your_api_key",
    "intensity": "0.8",
    "shape": "circle",
    "color": "#000000"
}

res = requests.post(url, files=files, data=data, stream=True)

if res.headers.get("Content-Type", "").startswith("image/"):
    with open("output.jpg", "wb") as f:
        for chunk in res.iter_content(1024):
            f.write(chunk)
else:
    print(res.json())  # error JSON
```

**Usage PHP**
```php
<?php
$url = "https://discardapi.dpdns.org/api/image/vignette";

$postfields = [
  "apikey" => "your_api_key",
  "file" => new CURLFile("input.jpg"),
  "intensity" => "0.8",
  "shape" => "circle",
  "color" => "#000000"
];

$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $url);
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, $postfields);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

$response = curl_exec($ch);
$contentType = curl_getinfo($ch, CURLINFO_CONTENT_TYPE);
curl_close($ch);

if (strpos($contentType, "image/") === 0) {
    file_put_contents("output.jpg", $response);
} else {
    echo $response; // JSON error
}
?>
```

### 📦 Catbox File Upload API (🟣POST)

**Endpoint:**
```nginx
POST https://discardapi.dpdns.org/api/catbox
```
**Parameters**

| Name   | Type   | Required | Description                         |
| ------ | ------ | -------- | ----------------------------------- |
| apikey | string | ✅        | Your API key                        |
| file   | file   | ✅        | File to upload (any type supported) |

**Response Format**
```json
{
  "status": true,
  "creator": "Qasim Ali 🩷",
  "result": {
    "url": "https://files.catbox.moe/abcd123.zip"
}
```

**Usage cURL**

```bash

curl -X POST "https://discardapi.dpdns.org/api/catbox" \
  -F "apikey=your_api_key" \
  -F "file=@example.zip"
```

**Usage JavaScript**
```js
import axios from "axios";
import FormData from "form-data";
import fs from "fs";

const form = new FormData();
form.append("apikey", "your_api_key");
form.append("file", fs.createReadStream("example.zip"));

const res = await axios.post("https://discardapi.dpdns.org/api/catbox", form, {
  headers: form.getHeaders()
});

console.log(res.data);
```

**Usage Python**

```py
import requests

url = "https://discardapi.dpdns.org/api/catbox"
files = {"file": open("example.pdf", "rb")}
data = {"apikey": "your_api_key"}

res = requests.post(url, files=files, data=data)
print(res.json())
```

**Usage PHP**

```php

<?php
$ch = curl_init("https://discardapi.dpdns.org/api/catbox");

$post = [
    "apikey" => "your_api_key",
    "file" => new CURLFile("example.docx")
];

curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $post);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

$response = curl_exec($ch);
curl_close($ch);

echo $response;
?>
```

**Usage Go**

```go

package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	file, _ := os.Open("example.mp4")
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("apikey", "your_api_key")
	part, _ := writer.CreateFormFile("file", "example.mp4")
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", "https://discardapi.dpdns.org/api/catbox", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
```

**Usage Java**

```java

import okhttp3.*;

import java.io.File;
import java.io.IOException;

public class Upload {
    public static void main(String[] args) throws IOException {
        OkHttpClient client = new OkHttpClient();

        File file = new File("example.mp3");

        RequestBody body = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("apikey", "your_api_key")
                .addFormDataPart("file", file.getName(),
                        RequestBody.create(file, MediaType.parse("application/octet-stream")))
                .build();

        Request request = new Request.Builder()
                .url("https://discardapi.dpdns.org/api/catbox")
                .post(body)
                .build();

        Response response = client.newCall(request).execute();
        System.out.println(response.body().string());
    }
}
```

**Usage C#**

```csharp

using System;
using System.Net.Http;
using System.Threading.Tasks;

class Program {
    static async Task Main() {
        var client = new HttpClient();
        var form = new MultipartFormDataContent();
        form.Add(new StringContent("your_api_key"), "apikey");
        form.Add(new StreamContent(System.IO.File.OpenRead("example.txt")), "file", "example.txt");

        var res = await client.PostAsync("https://discardapi.dpdns.org/api/catbox", form);
        var response = await res.Content.ReadAsStringAsync();
        Console.WriteLine(response);
    }
}
```

**Usage Swift**

```swift

import Foundation

let url = URL(string: "https://discardapi.dpdns.org/api/catbox")!
var request = URLRequest(url: url)
request.httpMethod = "POST"

let boundary = UUID().uuidString
request.setValue("multipart/form-data; boundary=\(boundary)", forHTTPHeaderField: "Content-Type")

var body = Data()
body.append("--\(boundary)\r\n".data(using: .utf8)!)
body.append("Content-Disposition: form-data; name=\"apikey\"\r\n\r\n".data(using: .utf8)!)
body.append("your_api_key\r\n".data(using: .utf8)!)

let fileURL = URL(fileURLWithPath: "example.mov")
let fileData = try Data(contentsOf: fileURL)

body.append("--\(boundary)\r\n".data(using: .utf8)!)
body.append("Content-Disposition: form-data; name=\"file\"; filename=\"example.mov\"\r\n".data(using: .utf8)!)
body.append("Content-Type: application/octet-stream\r\n\r\n".data(using: .utf8)!)
body.append(fileData)
body.append("\r\n".data(using: .utf8)!)
body.append("--\(boundary)--\r\n".data(using: .utf8)!)

request.httpBody = body

URLSession.shared.dataTask(with: request) { data, _, _ in
    if let data = data {
        print(String(data: data, encoding: .utf8)!)
    }
}.resume()
```

**Usage Kotlin**

```kotlin
import okhttp3.*
import java.io.File

fun main() {
    val client = OkHttpClient()
    val file = File("example.apk")

    val body = MultipartBody.Builder()
        .setType(MultipartBody.FORM)
        .addFormDataPart("apikey", "your_api_key")
        .addFormDataPart("file", file.name,
            file.asRequestBody("application/octet-stream".toMediaType()))
        .build()

    val request = Request.Builder()
        .url("https://discardapi.dpdns.org/api/catbox")
        .post(body)
        .build()

    client.newCall(request).execute().use { response ->
        println(response.body?.string())
    }
}
```

  
### 🚨 Common Error Codes

- ![](https://img.shields.io/badge/401-red) Please provide the apikey or Invalid apikey  
- ![](https://img.shields.io/badge/403-orange) This endpoint requires Special API key  
- ![](https://img.shields.io/badge/400-orange) Please provide the (query) e.g url  
- ![](https://img.shields.io/badge/429-yellow) Rate limit exceeded, please wait until reset



### 🚨 Error Handling

Comprehensive error codes with detailed descriptions and resolution steps.  
All errors include helpful context and suggestions.  

---

#### 🔑 Authentication Errors  

- ![](https://img.shields.io/badge/401-red) **notparam** → API key parameter missing  
  ➝ Add `?apikey=YOUR_KEY` to your request  

- ![](https://img.shields.io/badge/401-orange) **invalidKey** → Invalid API key provided  
  ➝ Check your key or generate a new one  

- ![](https://img.shields.io/badge/400-orange) **notquery** → Missing required query parameter  
  ➝ Please provide query (e.g. `url`)  

- ![](https://img.shields.io/badge/403-purple) **notspecial** → Requires premium API key  
  ➝ Upgrade your account for access  

---

#### ⏳ Rate Limiting & System  

- ![](https://img.shields.io/badge/429-yellow) **limit** → Rate limit exceeded  
  ➝ Wait for reset or upgrade plan  

- ![](https://img.shields.io/badge/500-gray) **internal** → Internal server error  
  ➝ Temporary issue, try again later  

- ![](https://img.shields.io/badge/404-blue) **notfound** → Endpoint or resource not found  
  ➝ Check the URL and parameters  

---

#### 💡 Best Practices  

- ✅ Always check the `status` field before processing results  
- ✅ Implement exponential backoff for **rate limit** errors  
- ✅ Log error codes for debugging & monitoring  
- ✅ Handle network timeouts gracefully



## 🎧 Support & Contact

Need help or want to upgrade? Our team is here to assist you with integration, troubleshooting, and custom solutions.

### 📞 Contact Methods

- ![](https://img.shields.io/badge/Email-blue) **Email Support** → [discardapi@gmail.com](mailto:discardapi@gmail.com?subject=Support&body=Hello%20Team,)  
- ![](https://img.shields.io/badge/WhatsApp-green) **Live Chat** → [Chat on WhatsApp](https://wa.me/923051391007?text=Hello%20I%20need%20support)  
- ![](https://img.shields.io/badge/Discord-purple) **Community Support** → [Join our Discord Server](https://discord.gg/YBkzCWqz)  
- ![](https://img.shields.io/badge/GitHub-black) **Documentation** → [GitHub Examples](https://github.com/GlobalTechInfo/discard-api)  

---

### ⏱ Response Times
- ✅ **Premium Support** → < 1 hour  
- 🆓 **Free Support** → < 24 hours  

---
© 2025 **Discard API** — Built with Go & Fiber  
