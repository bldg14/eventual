import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";
import { get } from "./http/get";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

// TODO: Remove this when wiring up the call to display the events in the ui.
// This only makes the api request and logs it to the console.
(async () => {
  try {
    let result = await get("/api/v1/events");
    console.log(result);
  } catch (error) {
    console.error(error);
  }
})();
