console.log("Mini Postman frontend loaded 🚀");

const form = document.getElementById("request-form");
const responseEl = document.getElementById("response");

form.addEventListener("submit", async (e) => {
  e.preventDefault();

  const method = document.getElementById("method").value;
  const url = document.getElementById("url").value;
  const headersRaw = document.getElementById("headers").value;
  const bodyRaw = document.getElementById("body").value;

  let headers = {};
  if (headersRaw.trim() !== "") {
    try {
      headers = JSON.parse(headersRaw);
    } catch {
      responseEl.innerHTML = `<span class="error">Invalid headers JSON</span>`;
      return;
    }
  }

  let body = "";
  if (
    ["POST", "PUT", "PATCH"].includes(method.toUpperCase()) &&
    bodyRaw.trim() !== ""
  ) {
    body = bodyRaw;
  }

  const payload = { method, url, headers, body };

  try {
    const res = await fetch("/send-request", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });

    const data = await res.json();
    responseEl.innerHTML = syntaxHighlight(data);
  } catch (err) {
    responseEl.innerHTML = `<span class="error">Error: ${err}</span>`;
  }
});

function syntaxHighlight(json) {
  if (typeof json != "string") {
    json = JSON.stringify(json, undefined, 2);
  }
  json = json
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;");
  return json.replace(
    /("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g,
    (match) => {
      let cls = "number";
      if (/^"/.test(match)) {
        if (/:$/.test(match)) {
          cls = "key";
        } else {
          cls = "string";
        }
      } else if (/true|false/.test(match)) {
        cls = "boolean";
      } else if (/null/.test(match)) {
        cls = "null";
      }
      return `<span class="${cls}">${match}</span>`;
    }
  );
}
