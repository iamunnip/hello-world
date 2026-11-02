from fastapi import FastAPI
from fastapi.responses import HTMLResponse

# Create the web app.
app = FastAPI()


# When someone opens the home page ("/"), run this function.
@app.get("/")
def hello():
    # Send this text back to the browser as a web page.
    return HTMLResponse("<h1>Hello, World from Python!</h1>")
