import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Pokemon } from "../../app/models/Pokemon";

const initialState: { [id: string]: Pokemon } = {};

export const pokemonDetailsSlice = createSlice({
  name: "pokemonsDetails",
  initialState,
  reducers: {
    getPokemonsDetail: (state:{ [id: string]: Pokemon } , action: PayloadAction<Pok>) => {},
  },
});

// TODO: Create the method to get all the details of
// export const fetchPokemonDetail = (id: number, ) => {
//   return async (dispatch, getState, extraArgument) => {
//     const s = getState();
//   // logica para traer el detalle y mostrarlo
//   };
// };

export const { getPokemonDetail } = pokemonDetailsSlice.actions;
