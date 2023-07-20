import React from "react";
import "./App.css";
import { Event } from "./event/event";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

function App() {
  return (
    <Container>
      <Row>
        <header>
          <h1>Eventual</h1>
        </header>
      </Row>
      <Row>
        <Event
          title="Bike Ride"
          description="Nice ride along the coast"
          start={new Date()}
          end={new Date()}
          url="https://www.google.com"
          email="hello@world.com"
        />
      </Row>
    </Container>
  );
}

export default App;
