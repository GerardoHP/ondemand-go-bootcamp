import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap-icons/font/bootstrap-icons.css';
import './Pokemon.css';
import { useState, useEffect } from "react";
import { Pokemon } from './Pokemon';

export function PokemonList({ pokemons }: any) {

    const [selected, setSelected] = useState(-1);
    const [selectedPokemon, setSelectedPokemon] = useState<{ id: number, name: string, image: string }>({ id: 0, name: '', image: '' });
    const [filterdPokemons, setFilteredPokemons] = useState(pokemons);

    const selectPokemon = (newId: number) => {
        setSelected(newId);
        const newSelectedPokemon = filterdPokemons.find(({ id }: { id: number }) => id === newId);
        setSelectedPokemon(newSelectedPokemon);
    };

    const searchFilter = ({ target }: { target: { value: string } }) => {
        const newFilteredPokemons = pokemons.filter((pok: { name: string }) => {            
            return pok.name.toLocaleLowerCase().includes(target.value.toLocaleLowerCase());
        });

        setFilteredPokemons(newFilteredPokemons);
    }

    useEffect(() => {
        if (filterdPokemons.length === 0) {
            setFilteredPokemons(pokemons);
        }
    }, pokemons)

    return (<div className='row' style={{ height: '100vh' }}>
        <div className="col-md-2 offset-md-1 h-100 d-inline-block PokemonContainer" >
            <div className='input-group mb-3 h-5'>
                <span className='input-group-text' id='search-button'><i className="bi bi-search"></i></span>
                <input type='text' className='form-control' placeholder='Search pokemon...' aria-label='Search pokemon' aria-describedby='search-button' onChange={searchFilter} />
            </div>
            <div className='overflow-auto h-90'>
                <div className='list-group'>
                    {filterdPokemons.map(({ id, name, image }: { id: number, name: string, image: string }) => {
                        return <button className={`list-group-item list-group-item-action ${id === selected ? 'active' : ''}`} onClick={() => selectPokemon(id)} key={id}>{name}</button>
                    })}
                </div>
            </div>
        </div>
        <div className='col-md-8'>
            <Pokemon {...selectedPokemon}></Pokemon>
        </div>
    </div>);
}
