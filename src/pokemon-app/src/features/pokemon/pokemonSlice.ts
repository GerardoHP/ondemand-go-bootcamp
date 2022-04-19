import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = {
  pokemons: [],
  status: "idle",
  error: null,
};

export const pokemonSlice = createSlice({
  name: "pokemons",
  initialState,
  reducers: {
    getPokemons: (
      state: {
        pokemons: { id: number; name: string; image: string }[];
      },
      action: PayloadAction<{ id: number; name: string; image: string }[]>
    ) => {
      action.payload.map((pokemon: any) => {
        const { ID, Name, Image } = pokemon;
        state.pokemons.push({ id: ID, name: Name, image: Image });
      });
    },
    addPokemon: (state: {
      pokemons: { id: any; name: string; image: string }[];
    }) => {
      state.pokemons.push({
        id: state.pokemons.length,
        name: `Pokemon ${state.pokemons.length}`,
        image: "asdf",
      });
    },
  },
});

export const fetchAllPokemons = () => {
  return async (
    dispatch: (arg0: any) => void,
    getState: any,
    extraArgument: { api: { url: string } }
  ) => {
    try {
      const s: any = getState();
      if (s && s.pokemon.pokemons.length === 0) {
        const pokemons = await axios.get(extraArgument.api.url);
        dispatch(getPokemons(pokemons.data));
      }
    } catch (err) {
      console.log(err);
    }
  };
};

export const { addPokemon, getPokemons } = pokemonSlice.actions;

export const selectPokemonCount = (state: any): number =>
  state.pokemon.pokemons.length;
export const selectAllPokemons = (state: any) => state.pokemon.pokemons;

export default pokemonSlice.reducer;
