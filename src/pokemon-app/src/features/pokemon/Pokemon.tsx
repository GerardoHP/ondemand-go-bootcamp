import { Pokemon as PokemonInterface } from "../../app/models/Pokemon";
import { useDispatch, useSelector } from "react-redux";
import { fetchPokemonDetail, IPokemonDetailState } from "../pokemon/pokemonDetailSlice";
import { useEffect } from "react";
import { PokemonDetail } from "../../app/models/PokemonDetail";
import { Abilities } from "./Components";

export function Pokemon({ Id, Name, Image, Url }: PokemonInterface) {

    const dispatch = useDispatch();

    const pokemon: PokemonDetail = useSelector(({ pokemonDetail }: { pokemonDetail: IPokemonDetailState }) => pokemonDetail.pokemons[Name.toLocaleLowerCase()] ?? null);
    console.log(pokemon);
    useEffect(() => {
        if (!pokemon) {
            dispatch(fetchPokemonDetail(Name, Url));
        }
    }, [Name]);

    return (<div>
        {pokemon ?
            <div className="pokemon">
                <div className="px-3 py-3 pt-md-5 pb-md-4 mx-auto text-center">
                    <h1 className="title display-4">{Name || pokemon.name}</h1>
                    <div className="container row">
                        <img src={Image} alt={Name} className="img-thumbnail col-md-1" />
                        <p className="descrtiption col-md-11 lead">Pokem ipsum dolor sit amet Azumarill Bidoof Gible Hippowdon consectetur adipisicing elit Regirock. Body Slam Helix Fossil Yanma Onix Abomasnow Machop Pupitar. Hoenn Poison Sting Kricketot Bayleef Rising Badge Abra Sandslash. Duis aute irure dolor in reprehenderit in voluptate Silver Tangela Lopunny Mantine Cleffa Mantine. Kanto Timburr Claydol Reuniclus consectetur adipisicing elit Pansear Dialga.</p>
                    </div>
                </div>
                {/* <div className="input-group mb-3">
                    <div className="input-group-prepend">
                        <span className="input-group-text" id="basic-addon1">Abilities</span>
                    </div>
                    {pokemon.abilities.map(({ ability, is_hidden, slot }) => {
                        return !is_hidden || ability.name
                    })}
                </div> */}
                <div className="container">
                    <div className="card-deck mb-3 text-center">
                        <Abilities abilities={pokemon.abilities} />
                    </div>

                </div>

            </div>
            : <div>no hay nada para ti </div>}
    </div>
    );
}


