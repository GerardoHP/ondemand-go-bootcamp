import { configureStore } from "@reduxjs/toolkit";
import pokemontReducer from "../features/pokemon/pokemonSlice";
import pokemonDeatilReducer from "../features/pokemon/pokemonDetailSlice";

export const store = configureStore({
  reducer: {
    pokemon: pokemontReducer,
    pokemonDetail: pokemonDeatilReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      thunk: {
        extraArgument: {
          api: { url: "http://localhost:8080" },
        },
      },
    }),
  // Si existieran mas reducers irian aqui
});
