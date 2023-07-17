import type { Actions } from './$types';

export const actions = {

  
    default: async ({ event, response }) => {
        // TODO log the user in
        const formData = new FormData()
        formData.append('file', event.currentTarget.files[0] as Blob)
        try {
            const res = await fetch(' http://127.0.0.1:3000/upload', {
                method: 'POST',
                body: formData
            })
            response.status = 200
            response.body = {
                message: 'success'
            }
            console.log(res)
        }
        catch (err) {
            response.status = 500
            response.body = {
                message: 'error'
            }
        }

    }



} satisfies Actions;