import React from "react";
import "./App.css";
import Heading from "./heading/heading";
import Event from "./event/event";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Heading>Eventual</Heading>
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
