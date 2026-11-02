use axum::{Router, response::Html, routing::get};

// Send a web page back to the browser.
async fn hello() -> Html<&'static str> {
    Html("<h1>Hello, World from Rust!</h1>")
}

#[tokio::main]
async fn main() {
    // When someone opens the home page ("/"), run the hello function.
    let app = Router::new().route("/", get(hello));

    // Start the server. It keeps running until you stop it.
    // 0.0.0.0 means it also works inside Docker.
    let listener = tokio::net::TcpListener::bind("0.0.0.0:8002")
        .await
        .unwrap();
    println!("Open http://127.0.0.1:8002 in your browser");
    axum::serve(listener, app).await.unwrap();
}
