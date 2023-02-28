import React from 'react';
import './App.css';
import Heading from './heading/heading'

class App extends React.Component {

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <Heading title="Eventual" />
        </header>
      </div>
    );
  }
}

export default App;
