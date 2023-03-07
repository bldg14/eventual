import React from 'react';
import './App.css';
import Heading from './heading/heading'

function Event({title, description}: {title: string; description: string;}) {
  return (
    <div>
      <Heading title={ title } />
      <p>
        { description }
      </p>
    </div>
  );
}

function App() {
    return (
      <div className="App">
        <header className="App-header">
          <Heading title="Eventual" />
        </header>
        <Event title="Bike Ride" description="Nice ride along the coast" />
      </div>
    );
}

export default App;
