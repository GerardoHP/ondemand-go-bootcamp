import { useState } from "react";
import { ability } from "../../app/models/PokemonDetail";

export const Abilities = ({ abilities }: { abilities: ability[] }) => {

    const [displayedAbilities, setDisplayedAbilities] = useState(abilities.filter(a => !a.is_hidden));
    // const hidden = displayedAbilities.filter(a => a.is_hidden).length > 0;
    const [hidden, setHidden] = useState(displayedAbilities.filter(a => a.is_hidden).length > 0);

    const displayHiddens = () => {
        setDisplayedAbilities(hidden ? abilities.filter(a => !a.is_hidden) : abilities);
        setHidden(!hidden);
    };

    return (<div className="card mb-4 shadow-sm">
        <div className="card-header">
            <h4 className="my-0 font-weight-normal">Ability</h4>
        </div>
        <div className="card-body">
            <h1 className="card-title pricing-card-title">
                {`Abilit${displayedAbilities.length > 1 ? 'ies' : 'y'}`}
                <a href='#' onClick={displayHiddens}><small className="text-muted"><i className={`bi bi-eye${hidden ? '-slash' : ''}`}></i></small></a>

            </h1>
            <ul className="list-unstyled mt-3 mb-4">
                {displayedAbilities.map((a) => <li>Name: {a.ability.name}, Slot: {a.slot}</li>)}
            </ul>
        </div>
    </div>);
}

const Card = ({ header }: { header: string }) => {
    return (<div className="card mb-4 shadow-sm">
        <div className="card-header">
            <h4 className="my-0 font-weight-normal">{header}</h4>
        </div>
        <div className="card-body">
            <h1 className="card-title pricing-card-title">$0 <small className="text-muted">/ mo</small></h1>
            <ul className="list-unstyled mt-3 mb-4">
                <li>10 users included</li>
                <li>2 GB of storage</li>
                <li>Email support</li>
                <li>Help center access</li>
            </ul>
            <button type="button" className="btn btn-lg btn-block btn-outline-primary">Sign up for free</button>
        </div>
    </div>);
}