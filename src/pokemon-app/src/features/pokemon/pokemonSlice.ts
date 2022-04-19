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
      // state.pokemons = action.payload;
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
          // const { id, name, image, url } = pokemon;
          // state.pokemons.push({ id, name, image, url });
          return { Id: ID, Name, Url, Image };
        }
      );
    },
    addPokemon: (state: { pokemons: Pokemon[] }) => {
      state.pokemons.push({
        Id: state.pokemons.length,
        Name: `Pokemon ${state.pokemons.length}`,
        Image: "asdf",
        Url: "",
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
        dispatch(setPokemons(pokemons.data));
      }
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
