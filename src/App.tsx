import React from "react";
import "./App.css";
import { Event } from "./event/event";

function App() {
  return (
    <div className="App min-vh-100 d-flex flex-column align-items-center">
      <header className="App-header">
        <h1>Eventual</h1>
      </header>
      <Event
        title="Bike Ride"
        description="Nice ride along the coast"
        start={new Date()}
        end={new Date()}
        url="https://www.google.com"
        email="hello@world.com"
      />
    </div>
  );
}

export default App;
