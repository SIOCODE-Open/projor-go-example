/** The user accounts in the system */
export interface IUserAccount {
    /** Unique identifier of the UserAccount */
    id: string;
    /** The username of the user account */
    username: string;
    /** The email address of the user account */
    email: string;
    /** The flags of the user account */
    flags: string;
}

export interface IUserAccountClient {
    create(data: IUserAccount): Promise<void>;
    update(data: IUserAccount): Promise<void>;
    delete(data: IUserAccount): Promise<void>;
    getById(id: string): Promise<IUserAccount>;
    getAll(): Promise<IUserAccount[]>;
}

class UserAccountClientImpl implements IUserAccountClient {
    private baseUrl: string;

    constructor(baseUrl: string) {
        this.baseUrl = baseUrl;
    }

    async create(data: IUserAccount): Promise<void> {
        const response = await fetch(`${this.baseUrl}/user-account`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        if (!response.ok) {
            throw new Error(response.statusText);
        }
    }

    async update(data: IUserAccount): Promise<void> {
        const response = await fetch(`${this.baseUrl}/user-account/${data.id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        if (!response.ok) {
            throw new Error(response.statusText);
        }
    }

    async delete(data: IUserAccount): Promise<void> {
        const response = await fetch(`${this.baseUrl}/user-account/${data.id}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        if (!response.ok) {
            throw new Error(response.statusText);
        }
    }

    async getById(id: string): Promise<IUserAccount> {
        const response = await fetch(`${this.baseUrl}/user-account/${id}`);
        if (!response.ok) {
            throw new Error(response.statusText);
        }
        return response.json();
    }

    async getAll(): Promise<IUserAccount[]> {
        const response = await fetch(`${this.baseUrl}/user-account`);
        if (!response.ok) {
            throw new Error(response.statusText);
        }
        return response.json();
    }
}
/** The products in the system */
export interface IProduct {
    /** Unique identifier of the Product */
    id: string;
    /** The name of the product */
    name: string;
    /** The price of the product */
    price: number;
    /** The description of the product */
    description: string;
}

export interface IProductClient {
    create(data: IProduct): Promise<void>;
    update(data: IProduct): Promise<void>;
    delete(data: IProduct): Promise<void>;
    getById(id: string): Promise<IProduct>;
    getAll(): Promise<IProduct[]>;
}

class ProductClientImpl implements IProductClient {
    private baseUrl: string;

    constructor(baseUrl: string) {
        this.baseUrl = baseUrl;
    }

    async create(data: IProduct): Promise<void> {
        const response = await fetch(`${this.baseUrl}/product`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        if (!response.ok) {
            throw new Error(response.statusText);
        }
    }

    async update(data: IProduct): Promise<void> {
        const response = await fetch(`${this.baseUrl}/product/${data.id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        if (!response.ok) {
            throw new Error(response.statusText);
        }
    }

    async delete(data: IProduct): Promise<void> {
        const response = await fetch(`${this.baseUrl}/product/${data.id}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        if (!response.ok) {
            throw new Error(response.statusText);
        }
    }

    async getById(id: string): Promise<IProduct> {
        const response = await fetch(`${this.baseUrl}/product/${id}`);
        if (!response.ok) {
            throw new Error(response.statusText);
        }
        return response.json();
    }

    async getAll(): Promise<IProduct[]> {
        const response = await fetch(`${this.baseUrl}/product`);
        if (!response.ok) {
            throw new Error(response.statusText);
        }
        return response.json();
    }
}

export interface IAppClient {
    userAccount: IUserAccountClient;
    product: IProductClient;
}

export function createAppClient(baseUrl: string): IAppClient {
    return {
        userAccount: new UserAccountClientImpl(baseUrl),
        product: new ProductClientImpl(baseUrl),
    };
}