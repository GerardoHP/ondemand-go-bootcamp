import { createSelector, createSlice, PayloadAction } from "@reduxjs/toolkit";
import axios from "axios";
import { Pokemon } from "../../app/models/Pokemon";
import { PokemonDetail } from "../../app/models/PokemonDetail";

export interface IPokemonDetailState {
  pokemons: { [name: string]: PokemonDetail };
}

const initialState: IPokemonDetailState = { pokemons: {} };

export const pokemonDetailsSlice = createSlice({
  name: "pokemonsDetails",
  initialState,
  reducers: {
    setPokemonsDetail: (
      state: any,
      action: PayloadAction<PokemonDetail>
    ) => {
      state.pokemons[action.payload.name] = action.payload;
    },
  },
});

export const fetchPokemonDetail = (name: string, url: string) => {
  return async (dispatch: (arg0: any) => void, getState: any) => {
    try {
      const { pokemonDetail }: { pokemonDetail: IPokemonDetailState } = getState();
      if (pokemonDetail && url !== "" && !pokemonDetail.pokemons[name]) {
        const pokemonDetail = await axios.get<PokemonDetail>(url);
        dispatch(setPokemonsDetail(pokemonDetail.data));
      }
    } catch (err) {
      console.log(err);
    }
  };
};

export const { setPokemonsDetail } = pokemonDetailsSlice.actions;

export default pokemonDetailsSlice.reducer;
