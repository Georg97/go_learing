using System;
using System.Threading.Channels;
using System.Threading.Tasks;

class Program
{
    static int[] s = { 7, 2, 8, -9, 4, 0 };
    static Channel<int> channel = Channel.CreateBounded<int>(3);

    static async Task Main(string[] args)
    {
        var t1 = SumAsync(s.AsSpan(0, s.Length / 2));
        var t2 = SumAsync(s.AsSpan(s.Length / 2));

        await channel.Writer.WriteAsync(66);

        int x = await channel.Reader.ReadAsync();
        int y = await channel.Reader.ReadAsync();
        int z = await channel.Reader.ReadAsync();

        Console.WriteLine($"{x} {y} {x + y}");
        Console.WriteLine(z);

        await Task.WhenAll(t1, t2);
    }

    static async Task SumAsync(ReadOnlySpan<int> slice)
    {
        int sum = 0;
        foreach (var num in slice)
        {
            sum += num;
        }

        await channel.Writer.WriteAsync(sum);
    }
}

