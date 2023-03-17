import React from 'react';
import './App.css';
import Heading from './heading/heading'
import Event from './event/event'

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
