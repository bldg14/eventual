const baseUrl = process.env.REACT_APP_BACKEND_URL || "";

export const get = async (endpoint: string, queryParameters?: Record<string, string>): Promise<any> => {
    const queryString = queryParameters ? new URLSearchParams(queryParameters).toString() : '';
    const url = queryString ? `${baseUrl}${endpoint}?${queryString}` : `${baseUrl}${endpoint}`;

    const headers = {
        "Content-Type": "application/json",
    };

    const config: RequestInit = {
        method: "GET",
        headers,
    };

    let response;
    try {
        response = await fetch(url, config);
        if (!response.ok) {
            throw new Error(String(response.status));
        }
    } catch (error) {
        throw new Error(`Fetching ${url} failed: ${error}`);
    }

    return response.json();
};
