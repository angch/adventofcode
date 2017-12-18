defmodule D01b do
  @moduledoc """
  Documentation for D01b for Advent of Code.
  """

  def index do
    readFile("input.txt")
    |> main
  end

  @doc """
  Test cases (Pressume everything is in even)

  ## Examples
      iex> D01b.main("1212")
      6
      iex> D01b.main("1221")
      0
      iex> D01b.main("123123")
      12
      iex> D01b.main("12131415")
      4
  """
  def main(input) do
    output = input
             |> String.split("", trim: true)
             |> Enum.map(&String.to_integer/1)
             |> split
             |> match
             |> sum
    output
  end

  def readFile(filename) do
    {:ok, input} = File.read(filename)
    input
  end

  def split(lists) do
    len = round(length(lists)/2)
    Enum.split(lists, len)
  end

  def match({lists1, lists2}) do
    res = Enum.zip(lists1, lists2)
    |> Enum.map(fn {ww, vv} ->
      if ww == vv, do: vv, else: 0
    end)
    res
  end

  def sum(lists) do
    Enum.reduce(lists, fn x,y ->
      (x + y)
    end) * 2
  end
end

D01b.index
|> IO.inspect
