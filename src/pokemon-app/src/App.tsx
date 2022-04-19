import React, { useState, useEffect } from 'react';
import logo from './logo.svg';
import './App.css';
import { pokemonSlice, selectPokemonCount, addPokemon, selectAllPokemons, fetchAllPokemons } from './features/pokemon/pokemonSlice';
import { useDispatch, useSelector } from 'react-redux';
import { Pokemon } from './features/pokemon/Pokemon'
import { PokemonList } from './features/pokemon/PokemonList';

function App() {

  const [counter, setCounter] = useState(0)

  const pokemonCount: number = useSelector(selectPokemonCount);
  const pokemons:any[] = useSelector(selectAllPokemons);
  const pokemonsFilter = useSelector((state:
    {
      pokemon: {
        pokemons: { id: any; name: string; image: string }[]
      }
    }): number =>
    state.pokemon.pokemons.filter((p) => p.id === 1).length)
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
