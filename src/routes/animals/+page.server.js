
export async function load () {
  const limit = 20
  // const data = await products.find({}).limit(limit).toArray()

  const data = ['Cat', 'Dog', 'Bird']

  return {
    products: data,
    next: data.length,
  }
}
