export default (error: any, context: any) => {
  console.log(error)
  context.error({ statusCode: 304, message: 'Server error' })
}
