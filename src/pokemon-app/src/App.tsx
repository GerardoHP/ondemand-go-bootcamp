import { useState, useEffect } from 'react';
import logo from './logo.svg';
import './App.css';
import { pokemonSlice, selectPokemonCount, addPokemon, selectAllPokemons } from './features/pokemon/pokemonSlice';
import { useDispatch, useSelector } from 'react-redux';
import { Pokemon } from './app/models/Pokemon';

function App() {

  const [counter, setCounter] = useState(0)

  const pokemonCount = useSelector(selectPokemonCount);
  console.log(pokemonCount);
  const pokemons = useSelector(selectAllPokemons);
  console.log(pokemons);
  const dispatch = useDispatch();

  const increaseCounter = () => {
    let newCounter = counter + 1;
    setCounter(newCounter);

    dispatch(addPokemon())
  };

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

        <p>Hay {pokemonCount} pokemons</p>
        {/* <p>Hay {pokemons} pokemons</p> */}

        {pokemons.map((pok: any) => { return pokemon(pok) })}

        <p>Has dado {counter} veces click! </p>

        <button name="button" onClick={increaseCounter}>Boton</button>
      </header>
    </div>
  );
}



const pokemon = ({ id = 0, name = '' }) => {
  return (
    <div key={id}>
      <p>id:</p>
      <p>{id}</p>
      <p>name:</p>
      <p>{name}</p>
    </div>
  );
}

export default App;
