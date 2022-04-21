import { createSelector, createSlice, PayloadAction } from "@reduxjs/toolkit";
import axios from "axios";
import { Pokemon } from "../../app/models/Pokemon";
import { PokemonDetail } from "../../app/models/PokemonDetail";

interface IState {
  pokemons: { [id: string]: PokemonDetail };
}

const initialState: IState = { pokemons: {} };

export const pokemonDetailsSlice = createSlice({
  name: "pokemonsDetails",
  initialState,
  reducers: {
    setPokemonsDetail: (
      state: IState,
      action: PayloadAction<PokemonDetail>
    ) => {
      state[action.payload.id] = action.payload;
    },
  },
});

export const fetchPokemonDetail = (id: number, url: string) => {
  return async (dispatch: (arg0: any) => void, getState: any) => {
    try {
      const s: { [id: string]: PokemonDetail } = getState();
      if (s && !s[id]) {
        const pokemonDetail = await axios.get<PokemonDetail>(url);
        dispatch(setPokemonsDetail(pokemonDetail.data));
      }
    } catch (err) {
      console.log(err);
    }
  };
};

export const { setPokemonsDetail } = pokemonDetailsSlice.actions;

const selectById = createSelector(
  [(state) => state, (state, id) => id],
  (state, id) => state[id]
);

export default pokemonDetailsSlice.reducer;
