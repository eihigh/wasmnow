# wasmnow
wasmnow is a simple tool that helps you create Go programs that run in your web browser.

## How to Use
Run `go run github.com/eihigh/wasmnow@latest` and it will generate all the necessary files for browser execution in a `wasmnow` directory. A server will automatically start at `http://localhost:8080/` - just open this URL in your browser to test your program.

Once generated, you're free to use the `wasmnow` contents however you like: distribute them on your webpage, package them into a ZIP file for uploading to browser game platforms like itch.io, customize the `index.html` design, and more.

## Running the Demo
A sample program is available in the example directory. Follow these steps to try it out:

1. `cd example/`
2. Start wasmnow by running `go run ..` (dot-dot)
3. Open `http://localhost:8080` in your browser
4. If you see `Hello, WebAssembly!`, everything is working correctly
