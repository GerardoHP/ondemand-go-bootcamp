import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import axios from "axios";
import { Pokemon } from "../../app/models/Pokemon";

interface IState {
  pokemons: Pokemon[];
  status: string;
  error: string | null;
}

const initialState: IState = {
  pokemons: [],
  status: "idle",
  error: null,
};

export const pokemonSlice = createSlice({
  name: "pokemons",
  initialState,
  reducers: {
    setPokemons: (state: IState, action: PayloadAction<any[]>) => {
      state.pokemons = action.payload.map(
        ({
          ID,
          Name,
          Url,
          Image,
        }: {
          ID: number;
          Name: string;
          Url: string;
          Image: string;
        }) => {
          return { Id: ID, Name, Url, Image };
        }
      );
    },
    addPokemon: (state: IState, action: PayloadAction<any>) => {
      const pokemon: Pokemon = { ...action.payload };
      pokemon.Id = action.payload.ID;
      state.pokemons.push(pokemon);
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
        const pokemons = await axios.get(`${extraArgument.api.url}/pokemons`);
        dispatch(setPokemons(pokemons.data));
      }
    } catch (err) {
      console.log(err);
    }
  };
};

export const getPokemon = (pokemonName: string) => {
  return async (
    dispatch: (arg0: any) => void,
    getState: any,
    extraArgument: { api: { url: string } }
  ) => {
    try {
      const pokemon = await axios.get(
        `${extraArgument.api.url}/pokemon/${pokemonName}`
      );

      if (pokemon.status !== 200) {
        throw new Error("pokemon not found");
      }

      dispatch(addPokemon(pokemon.data));
    } catch (err) {
      console.log(err);
    }
  };
};

export const { addPokemon, setPokemons } = pokemonSlice.actions;

export const selectPokemonCount = (state: any): number =>
  state.pokemon.pokemons.length;
export const selectAllPokemons = (state: any) => state.pokemon.pokemons;

export default pokemonSlice.reducer;
