import { configureStore } from '@reduxjs/toolkit'
import pokemontReducer from "../features/pokemon/pokemonSlice";

export const store = configureStore({
    reducer: {
        counter: pokemontReducer
    },
    // Si existieran mas reducers irian aqui
});