import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  pokemonsDetails: [],
};

export const pokemonDetailsSlice = createSlice({
  name: "pokemonsDetails",
  initialState,
  reducers: {
    getPokemonDetail: () => {},
  },
});


// TODO: Create the method to get all the details of 
// export const fetchPokemonDetail = (id: number, ) => {
//   return async (dispatch, getState, extraArgument) => {
//     const s = getState();
//   };
// };

export const { getPokemonDetail } = pokemonDetailsSlice.actions;
