import { useEffect, useState } from 'react';
import {
    IUserAccount,
    IProduct,
    IAppClient,
    createAppClient
} from "./client";

const appClient = createAppClient("http://localhost:8080");

function UserAccountTable(
    props: {
        value: Array<IUserAccount>
    }
) {
    const tableRows = props.value.map(
        (data) => (
            <tr key={data.id}>
                <td>{data.id}</td>
                <td>{ data.username }</td>
                <td>{ data.email }</td>
                <td>{ data.flags }</td>
            </tr>
        )
    );
    return <table border="1">
        <thead>
            <tr>
                <th>Id</th>
                <th>Username</th>
                <th>Email</th>
                <th>Flags</th>
            </tr>
        </thead>
        <tbody>
            {tableRows}
        </tbody>
    </table>;
}

function UserAccountCreateForm(
    props: {
        onCreated: () => void
    }
) {
    const [id, setId] = useState("");
    const [username, setUsername] = useState(
        ""
    );
    const [email, setEmail] = useState(
        ""
    );
    const [flags, setFlags] = useState(
        ""
    );

    const clearForm = () => {
        setId("");
        setUsername(
            ""
        );
        setEmail(
            ""
        );
        setFlags(
            ""
        );
    }

    const onSave = async () => {
        try {
            await appClient.userAccount.create({
                id,
                username,
                email,
                flags: flags.split(","),
            });
            props.onCreated();
            clearForm();
        } catch (error) {
            console.error(error);
        }
    };

    return <form>
        <div>
            <label>
                Id
            </label>
            <input
                type="text"
                value={id}
                onChange={(e) => setId(e.target.value)}
            />
        </div>
        <div>
            <label>
                Username
            </label>
            <input
                type="text"
                value={ username }
                onChange={(e) => setUsername(e.target.value)}
            />
        </div>
        <div>
            <label>
                Email
            </label>
            <input
                type="text"
                value={ email }
                onChange={(e) => setEmail(e.target.value)}
            />
        </div>
        <div>
            <label>
                Flags
            </label>
            <input
                type="text"
                value={ flags }
                onChange={(e) => setFlags(e.target.value)}
            />
        </div>
        <button type="button" onClick={onSave}>Save</button>
    </form>;
}
function ProductTable(
    props: {
        value: Array<IProduct>
    }
) {
    const tableRows = props.value.map(
        (data) => (
            <tr key={data.id}>
                <td>{data.id}</td>
                <td>{ data.name }</td>
                <td>{ data.price }</td>
                <td>{ data.description }</td>
            </tr>
        )
    );
    return <table border="1">
        <thead>
            <tr>
                <th>Id</th>
                <th>Name</th>
                <th>Price</th>
                <th>Description</th>
            </tr>
        </thead>
        <tbody>
            {tableRows}
        </tbody>
    </table>;
}

function ProductCreateForm(
    props: {
        onCreated: () => void
    }
) {
    const [id, setId] = useState("");
    const [name, setName] = useState(
        ""
    );
    const [price, setPrice] = useState(
        0.0
    );
    const [description, setDescription] = useState(
        ""
    );

    const clearForm = () => {
        setId("");
        setName(
            ""
        );
        setPrice(
            0.0
        );
        setDescription(
            ""
        );
    }

    const onSave = async () => {
        try {
            await appClient.product.create({
                id,
                name,
                price,
                description,
            });
            props.onCreated();
            clearForm();
        } catch (error) {
            console.error(error);
        }
    };

    return <form>
        <div>
            <label>
                Id
            </label>
            <input
                type="text"
                value={id}
                onChange={(e) => setId(e.target.value)}
            />
        </div>
        <div>
            <label>
                Name
            </label>
            <input
                type="text"
                value={ name }
                onChange={(e) => setName(e.target.value)}
            />
        </div>
        <div>
            <label>
                Price
            </label>
            <input
                type="text"
                value={ price }
                onChange={(e) => setPrice(e.target.value)}
            />
        </div>
        <div>
            <label>
                Description
            </label>
            <input
                type="text"
                value={ description }
                onChange={(e) => setDescription(e.target.value)}
            />
        </div>
        <button type="button" onClick={onSave}>Save</button>
    </form>;
}

export function App() {
    const [userAccounts, setUserAccounts] = useState<Array<IUserAccount>>([]);
    const [products, setProducts] = useState<Array<IProduct>>([]);

    const populateUserAccounts = async () => {
        try {
            const data = await appClient.userAccount.getAll();
            setUserAccounts(data || []);
        } catch (error) {
            console.error(error);
        }
    };
    const populateProducts = async () => {
        try {
            const data = await appClient.product.getAll();
            setProducts(data || []);
        } catch (error) {
            console.error(error);
        }
    };

    useEffect(() => {
        populateUserAccounts();
        populateProducts();
    }, []);

    return (
        <div>
            <h2>UserAccount</h2>
            <UserAccountTable value={ userAccounts } />
            <UserAccountCreateForm onCreated={populateUserAccounts} />
            <h2>Product</h2>
            <ProductTable value={ products } />
            <ProductCreateForm onCreated={populateProducts} />
        </div>
    );
}