namespace SkyConvert.Api.Common;

internal static class FunctionalExtensions
{
    public static TOut Pipe<TIn, TOut>(this TIn input, Func<TIn, TOut> func) => func(input);
}
