import { useState, useEffect } from 'react';
import logo from './logo.svg';
import './App.css';

function App() {

  const [counter, setCounter] = useState(0)

  const increaseCounter = () => {
    let newCounter = counter + 1;
    setCounter(newCounter)
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>

        <p>Has dado {counter} veces click! </p>

        <button name="button" onClick={increaseCounter}>Boton</button>
      </header>
    </div>
  );
}

export default App;
