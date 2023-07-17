export const load = async ({ params }) => {
    const res = await fetch('http://127.0.0.1:3000', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
    });

    const data = await res.json();
    
    
    return {
        post: {
            title: 'My Post',
            content: `This is my post content ${JSON.stringify(data)}` 
        }
    };
};
