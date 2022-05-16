import React, { useEffect } from 'react';
import './App.css';
import { selectPokemonCount, selectAllPokemons, fetchAllPokemons } from './features/pokemon/pokemonSlice';
import { useDispatch, useSelector } from 'react-redux';
import { PokemonList } from './features/pokemon/PokemonList';

function App() {

  const pokemonCount: number = useSelector(selectPokemonCount);
  const pokemons: any[] = useSelector(selectAllPokemons);
  const dispatch = useDispatch();

  useEffect(() => {
    if (pokemonCount === 0) {
      dispatch(fetchAllPokemons());
    }
  });

  return (
    <div className="App">
      <PokemonList pokemons={pokemons}></PokemonList>
    </div>
  );
}

export default App;
