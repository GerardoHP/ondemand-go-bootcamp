import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";
import { Pokemon } from "../../app/models/Pokemon";

// const initialState = { value: [new Pokemon(0, "primis", "asdf")] };
const initialState = {
    pokemons: [
        { id: 0, name: "primis", image: "asdf" },
        { id: 1, name: "secus", image: "asdf" },
    ],
    status: 'idle',
    error: null,
};

export const pokemonSlice = createSlice({
    name: 'pokemons',
    initialState,
    reducers: {
        addPokemon: state => {
            // state.value.pokemons = [...state.value.pokemons, new Pokemon(state.value.pokemons.length, `Pokemon ${state.value.pokemons.length}`, "asdf")];
            state.pokemons.push({ id: state.pokemons.length, name: `Pokemon ${state.pokemons.length}`, image: "asdf" });
        },
    }
});

export const { addPokemon } = pokemonSlice.actions;

export const selectPokemonCount = (state: any): number => state.counter.pokemons.length;
export const selectAllPokemons = (state: any) => state.counter.pokemons;

export default pokemonSlice.reducer;

// const fetchPokemons = ()=>{
//     return async(dispatch, getState)=>{
//         try{
//             const pokemons = await axios.get("http://localhost:8080/pokemons");
//             dispatch()
//         }
//     }
// }